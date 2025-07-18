package valetransporte

import (
	"lapasta/internal/models"
)

type ValeService interface {
	ListarVales(page int) ([]models.ValeTransportePagamento, error)
	ListarValesDaSemana(semana int, page int, mes int) ([]models.ValeTransportePagamento, error)
	AtualizarStatusVale(id int, statusIdVale int) error
}

type valeService struct {
	repo ValeRepository
}

func NovoValeService(repo ValeRepository) ValeService {
	return &valeService{
		repo: repo,
	}
}

func (s *valeService) ListarVales(page int) ([]models.ValeTransportePagamento, error) {
	return s.repo.ListarVales(page)
}

func (s *valeService) ListarValesDaSemana(semana int, page int, mes int) ([]models.ValeTransportePagamento, error) {
	return s.repo.ListarValesDaSemana(semana, page, mes)
}

func (s *valeService) AtualizarStatusVale(id int, statusIdVale int) error {
	return s.repo.AtualizarStatusVale(id, statusIdVale)
}
