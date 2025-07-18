package sql

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"lapasta/internal/models"
	"time"
)

func (r *SQLStr) CriarFuncionario(funcionario *models.Funcionario) error {
	var existe bool
	checkQuery := `SELECT 1 FROM Funcionarios WHERE Cpf = @Cpf`
	err := r.db.QueryRow(checkQuery, sql.Named("Cpf", funcionario.Cpf)).Scan(&existe)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("erro ao verificar CPF existente: %w", err)
	}
	if existe {
		return fmt.Errorf("já existe um funcionário com o CPF %s", funcionario.Cpf)
	}

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}
	hashedPassword := sha256.Sum256([]byte(funcionario.Senha))

	funcQuery := `
    INSERT INTO Funcionarios (
        Nome, Sobrenome, Cpf, Rg, DataNasc, Email, Senha, Cargo, 
        DateAdmissao, HoraEntrada, HoraSaida, Salario, Admin,
        ValeTransporteSemanal, Status, ChavePix
    ) 
    VALUES (
        @Nome, @Sobrenome, @Cpf, @Rg, @DataNasc, @Email, @Senha, @Cargo, 
        @DateAdmissao, @HoraEntrada, @HoraSaida, @Salario, @Admin,
        @ValeTransporteSemanal, @Status, @ChavePix
    )
`

	_, err = tx.Exec(funcQuery,
		sql.Named("Nome", funcionario.Nome),
		sql.Named("Sobrenome", funcionario.Sobrenome),
		sql.Named("Cpf", funcionario.Cpf),
		sql.Named("Rg", funcionario.Rg),
		sql.Named("DataNasc", funcionario.DataNasc),
		sql.Named("Email", funcionario.Email),
		sql.Named("Senha", hashedPassword[:]),
		sql.Named("Cargo", funcionario.Cargo),
		sql.Named("DateAdmissao", funcionario.DateAdmissao),
		sql.Named("HoraEntrada", funcionario.HoraEntrada),
		sql.Named("HoraSaida", funcionario.HoraSaida),
		sql.Named("Salario", funcionario.Salario),
		sql.Named("Admin", funcionario.Admin),
		sql.Named("ValeTransporteSemanal", funcionario.ValeTransporteSemanal),
		sql.Named("Status", 1),
		sql.Named("ChavePix", funcionario.ChavePix),
	)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao inserir funcionário: %w", err)
	}

	authQuery := `
		INSERT INTO AUTH (username, password)
		VALUES (@Username, @Password)
	`

	_, err = tx.Exec(authQuery,
		sql.Named("Username", funcionario.Email),
		sql.Named("Password", hashedPassword[:]),
	)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao inserir credenciais de autenticação: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("erro ao finalizar transação: %w", err)
	}

	return nil
}

func (r *SQLStr) ListarFuncionarios(page int) ([]models.Funcionario, error) {
	var funcionarios []models.Funcionario

	limit := 10
	offset := (page - 1) * limit

	query := `
		SELECT 
			Id, Nome, Sobrenome, Cpf, Rg, DataNasc, Email, Cargo, 
			DateAdmissao, HoraEntrada, HoraSaida, ValeTransporteSemanal, Status, ChavePix, Salario, Admin
		FROM Funcionarios WITH (NOLOCK)
		ORDER BY DateAdmissao DESC
		OFFSET @Offset ROWS
		FETCH NEXT @Limit ROWS ONLY
	`

	rows, err := r.db.Query(query, sql.Named("Offset", offset), sql.Named("Limit", limit))
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar funcionários: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var f models.Funcionario
		if err := rows.Scan(
			&f.Id, &f.Nome, &f.Sobrenome, &f.Cpf, &f.Rg, &f.DataNasc, &f.Email, &f.Cargo,
			&f.DateAdmissao, &f.HoraEntrada, &f.HoraSaida, &f.ValeTransporteSemanal, &f.Status, &f.ChavePix, &f.Salario, &f.Admin,
		); err != nil {
			return nil, fmt.Errorf("erro ao escanear funcionário: %w", err)
		}
		funcionarios = append(funcionarios, f)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das linhas: %w", err)
	}

	return funcionarios, nil
}

func (r *SQLStr) BuscarFuncionarioPorCPF(cpf string) (models.FuncionarioComPontos, error) {
	var funcionario models.Funcionario
	query := "SELECT Id, Nome FROM Funcionarios WITH (NOLOCK) WHERE Cpf = @Cpf"

	err := r.db.QueryRow(query, sql.Named("Cpf", cpf)).Scan(&funcionario.Id, &funcionario.Nome)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.FuncionarioComPontos{}, fmt.Errorf("funcionário com CPF %s não encontrado", cpf)
		}
		return models.FuncionarioComPontos{}, fmt.Errorf("erro ao buscar funcionário: %w", err)
	}

	funcionarioComPontos := models.FuncionarioComPontos{
		Funcionario: funcionario,
	}

	dia := time.Now().Format("2006-01-02")
	ponto, err := r.ListarPontosPorIdEDia(funcionario.Id, dia)
	if err != nil {
		return models.FuncionarioComPontos{}, fmt.Errorf("erro ao buscar pontos para o funcionário %d no dia %s: %w", funcionario.Id, dia, err)
	}

	funcionarioComPontos.Pontos = append(funcionarioComPontos.Pontos, ponto)
	if ponto.HManha.Valid && ponto.HAlmocoSaida.Valid && ponto.HAlmocoRetorno.Valid && ponto.HNoite.Valid {
		funcionarioComPontos.BotaoStatus = "Todos os pontos registrados"
	} else if ponto.HManha.Valid && !ponto.HAlmocoSaida.Valid {
		funcionarioComPontos.BotaoStatus = "Pausa"
	} else if ponto.HAlmocoSaida.Valid && !ponto.HAlmocoRetorno.Valid {
		funcionarioComPontos.BotaoStatus = "Retorno"
	} else if ponto.HAlmocoRetorno.Valid && !ponto.HNoite.Valid {
		funcionarioComPontos.BotaoStatus = "Saida"
	} else {
		funcionarioComPontos.BotaoStatus = "Entrada"
	}
	return funcionarioComPontos, nil
}

func (r *SQLStr) BuscarFuncionarioPorID(id int) (*models.Funcionario, error) {
	var f models.Funcionario

	query := `
		SELECT 
			Id, Nome, Sobrenome, Cpf, Rg, DataNasc, Email, Cargo, 
			DateAdmissao, HoraEntrada, HoraSaida, ValeTransporteSemanal, Status, 
			ChavePix, Salario, Admin
		FROM Funcionarios WITH (NOLOCK)
		WHERE Id = @Id
	`

	err := r.db.QueryRow(query, sql.Named("Id", id)).Scan(
		&f.Id, &f.Nome, &f.Sobrenome, &f.Cpf, &f.Rg, &f.DataNasc, &f.Email, &f.Cargo,
		&f.DateAdmissao, &f.HoraEntrada, &f.HoraSaida, &f.ValeTransporteSemanal, &f.Status,
		&f.ChavePix, &f.Salario, &f.Admin,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // retorna nulo se não encontrar
		}
		return nil, fmt.Errorf("erro ao buscar funcionário por ID: %w", err)
	}

	return &f, nil
}
