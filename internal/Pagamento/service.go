package pagamento

import (
	"lapasta/internal/models"
)

type PagamentoService interface {
	CriarPagamento(pagamento *models.Pagamento) error
	ListarPagamentos(page int) ([]models.Pagamento, error)
	ListarPagamentosPorDia(date string, page int) ([]models.Pagamento, error)
	ListarPagamentosPorMes(month int, page int) ([]models.Pagamento, error)
	AtualizarStatusPagamento(id int, statusId int) error
}

type pagamentoService struct {
	repo PagamentoRepository
}

func NovoPagamentoService(repo PagamentoRepository) PagamentoService {
	return &pagamentoService{
		repo: repo,
	}
}

func (s *pagamentoService) CriarPagamento(pagamento *models.Pagamento) error {
	return s.repo.CriarPagamento(pagamento)
}

func (s *pagamentoService) ListarPagamentos(page int) ([]models.Pagamento, error) {
	return s.repo.ListarPagamentos(page)
}
func (s *pagamentoService) ListarPagamentosPorDia(date string, page int) ([]models.Pagamento, error) {
	return s.repo.ListarPagamentosPorDia(date, page)
}
func (s *pagamentoService) ListarPagamentosPorMes(month int, page int) ([]models.Pagamento, error) {
	return s.repo.ListarPagamentosPorMes(month, page)
}
func (s *pagamentoService) AtualizarStatusPagamento(id int, statusId int) error {
	return s.repo.AtualizarStatusPagamento(id, statusId)
}
