package sql

import (
	"database/sql"
	"errors"
	"lapasta/internal/models"
)

func (s *SQLStr) BuscarMotoristasAtivos() ([]models.Motorista, error) {
	rows, err := s.db.Query(`SELECT id, nome, cpf FROM Motoristas WHERE ativo = 1`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var motoristas []models.Motorista
	for rows.Next() {
		var m models.Motorista
		if err := rows.Scan(&m.Id, &m.Nome, &m.CPF); err != nil {
			return nil, err
		}
		motoristas = append(motoristas, m)
	}

	if len(motoristas) == 0 {
		return nil, errors.New("nenhum motorista ativo encontrado")
	}
	return motoristas, nil
}

func (s *SQLStr) SalvarEmissaoNota(nota models.EmissaoNota) error {
	_, err := s.db.Exec(`
		INSERT INTO EmissaoNotas (NumeroNota, Valor, DataEmissao, Descricao, MotoristaId)
		VALUES (@NumeroNota, @Valor, @DataEmissao, @Descricao, @MotoristaId)`,
		sql.Named("NumeroNota", nota.NumeroNota),
		sql.Named("Valor", nota.Valor),
		sql.Named("DataEmissao", nota.DataEmissao),
		sql.Named("Descricao", nota.Descricao),
		sql.Named("MotoristaId", nota.MotoristaId),
	)
	return err
}
