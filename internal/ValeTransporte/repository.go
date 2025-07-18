package valetransporte

import (
	database "lapasta/database"
	"lapasta/internal/models"
)

type ValeRepository interface {
	ListarVales(page int) ([]models.ValeTransportePagamento, error)
	ListarValesDaSemana(semana int, page int, mes int) ([]models.ValeTransportePagamento, error)
	AtualizarStatusVale(id int, statusIdVale int) error
}

type valeRepository struct {
	db *database.SQLStr
}

func NovoValeRepository(db *database.SQLStr) ValeRepository {
	return &valeRepository{
		db: db,
	}
}

func (r *valeRepository) ListarVales(page int) ([]models.ValeTransportePagamento, error) {
	return r.db.ListarVales(page)
}

func (r *valeRepository) ListarValesDaSemana(semana int, page int, mes int) ([]models.ValeTransportePagamento, error) {
	return r.db.ListarValesDaSemana(semana, page, mes)
}

func (r *valeRepository) AtualizarStatusVale(id int, statusIdVale int) error {
	return r.db.AtualizarStatusVale(id, statusIdVale)
}
