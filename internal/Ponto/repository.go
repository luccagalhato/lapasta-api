// repository/ponto.go
package ponto

import (
	database "lapasta/database"
	"lapasta/internal/models"
	"time"
)

type PontoRepository interface {
	ListarPontos(page int) ([]models.Ponto, error)
	ListarPontosPorId(idFuncionario int, page int) ([]models.Ponto, error)
	ListarPontosPorIdEDia(idFuncionario int, dia string) (models.Ponto, error)
	ListarPontosPorData(startDate, endDate time.Time, page int) ([]models.Ponto, error)
	ListarPontosPorDataId(IdFuncionario int, startDate time.Time, endDate time.Time, page int) ([]models.Ponto, error)
	RegistrarEntrada(idFuncionario int) error
	RegistrarSaidaAlmoco(idFuncionario int) error
	RegistrarRetornoAlmoco(idFuncionario int) error
	RegistrarSaida(idFuncionario int) error
	GerarRelatorioMensal(mes int, ano int, emailAdmin string) (string, error)
}

type pontoRepository struct {
	db *database.SQLStr
}

func NovoPontoRepository(db *database.SQLStr) PontoRepository {
	return &pontoRepository{
		db: db,
	}
}

func (r *pontoRepository) ListarPontos(page int) ([]models.Ponto, error) {
	return r.db.ListarPontos(page)
}

func (r *pontoRepository) ListarPontosPorId(idFuncionario int, page int) ([]models.Ponto, error) {
	return r.db.ListarPontosPorId(idFuncionario, page)
}

func (r *pontoRepository) ListarPontosPorIdEDia(idFuncionario int, dia string) (models.Ponto, error) {
	return r.db.ListarPontosPorIdEDia(idFuncionario, dia)
}

func (r *pontoRepository) RegistrarEntrada(idFuncionario int) error {
	return r.db.RegistrarEntrada(idFuncionario)
}

func (r *pontoRepository) ListarPontosPorData(startDate, endDate time.Time, page int) ([]models.Ponto, error) {
	return r.db.ListarPontosPorData(startDate, endDate, page)
}
func (r *pontoRepository) ListarPontosPorDataId(idFuncionario int, startDate time.Time, endDate time.Time, page int) ([]models.Ponto, error) {
	return r.db.ListarPontosPorDataId(idFuncionario, startDate, endDate, page)
}

func (r *pontoRepository) RegistrarSaidaAlmoco(idFuncionario int) error {
	return r.db.RegistrarSaidaAlmoco(idFuncionario)
}

func (r *pontoRepository) RegistrarRetornoAlmoco(idFuncionario int) error {
	return r.db.RegistrarRetornoAlmoco(idFuncionario)
}
	 	
func (r *pontoRepository) RegistrarSaida(idFuncionario int) error {
	return r.db.RegistrarSaida(idFuncionario)
}
func (r*pontoRepository)GerarRelatorioMensal(mes int, ano int, emailAdmin string) (string, error) {
	return r.db.GerarRelatorioMensal(mes, ano, emailAdmin)
}