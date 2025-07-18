package pagamento

import (
	database "lapasta/database"
	"lapasta/internal/models"
)

type PagamentoRepository interface {
	CriarPagamento(pagamento *models.Pagamento) error
	ListarPagamentos(page int) ([]models.Pagamento, error)
	ListarPagamentosPorMes(mes int, page int) ([]models.Pagamento, error)
	ListarPagamentosPorDia(date string, page int) ([]models.Pagamento, error)
	AtualizarStatusPagamento(id int, statusId int) error
}

type pagamentoRepository struct {
	db *database.SQLStr
}

func NovoPagamentoRepository(db *database.SQLStr) PagamentoRepository {
	return &pagamentoRepository{
		db: db,
	}
}

func (r *pagamentoRepository) CriarPagamento(pagamento *models.Pagamento) error {
	return r.db.CriarPagamento(pagamento)
}

func (r *pagamentoRepository) ListarPagamentos(page int) ([]models.Pagamento, error) {
	return r.db.ListarPagamentos(page)
}
func (r *pagamentoRepository) ListarPagamentosPorMes(mes int, page int) ([]models.Pagamento, error) {
	return r.db.ListarPagamentosPorMes(mes, page)
}
func (r *pagamentoRepository) ListarPagamentosPorDia(date string, page int) ([]models.Pagamento, error) {
	return r.db.ListarPagamentosPorDia(date, page)
}
func (r *pagamentoRepository) AtualizarStatusPagamento(id int, statusId int) error {
	return r.db.AtualizarStatusPagamento(id, statusId)
}
