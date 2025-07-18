package sql

import (
	"database/sql"
	"fmt"
	"lapasta/internal/models"
)

func (s *SQLStr) Autenticar(username string) (models.Login, error) {
	var Auth models.Login

	query := `SELECT username, password FROM AUTH WITH (NOLOCK) WHERE username = @Username`
	err := s.db.QueryRow(query, sql.Named("Username", username)).Scan(&Auth.Username, &Auth.PasswordCriptografado)
	if err != nil {
		if err == sql.ErrNoRows {
			return Auth, fmt.Errorf("usuário não encontrado")
		}
		return Auth, err
	}

	return Auth, nil
}

func (s *SQLStr) GetFuncionario(email string) (models.User, error) {
	var user models.User
	var admin int

	query := "SELECT ID, Email, Admin, Nome, Sobrenome FROM Funcionarios WITH (NOLOCK) WHERE Email = @Email"
	err := s.db.QueryRow(query, sql.Named("Email", email)).Scan(&user.ID, &user.Email, &admin, &user.Nome, &user.Sobrenome)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("usuário não encontrado")
		}
		return user, err
	}

	user.Admin = admin
	if admin == 1 {
		user.Tipo = "admin"
	} else {
		user.Tipo = "funcionario"
	}

	return user, nil
}

func (s *SQLStr) GetMotorista(email string) (models.User, error) {
	var user models.User

	query := `SELECT Id, Email, Nome FROM Motoristas WITH (NOLOCK) WHERE Email = @Email`
	err := s.db.QueryRow(query, sql.Named("Email", email)).Scan(&user.ID, &user.Email, &user.Nome)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("motorista não encontrado")
		}
		return user, err
	}

	user.Tipo = "motorista"
	return user, nil
}

func (s *SQLStr) GetUsuario(email string) (models.User, error) {
	user, err := s.GetFuncionario(email)
	if err == nil {
		return user, nil
	}

	user, err = s.GetMotorista(email)
	if err == nil {
		return user, nil
	}

	return models.User{}, fmt.Errorf("usuário não encontrado")
}
