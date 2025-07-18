package auth

import (
	database "lapasta/database"
	"lapasta/internal/models"
)

type AuthRepository interface {
	Autenticar(username string) (models.Login, error)
	GetUsuario(username string) (models.User, error)
}

type authRepository struct {
	db *database.SQLStr
}

func NewAuthRepository(db *database.SQLStr) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Autenticar(username string) (models.Login, error) {
	login, err := r.db.Autenticar(username)
	if err != nil {
		return models.Login{}, err
	}
	return login, nil
}

func (r *authRepository) GetUsuario(username string) (models.User, error) {
	user, err := r.db.GetUsuario(username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
