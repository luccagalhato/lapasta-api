package pedidofornecedor

import "lapasta/internal/models"

type PedidoFornecedorService interface {
	CriarPedidoFornecedor(pedido *models.PedidoFornecedor) error
	ListarPedidosPorFornecedor(fornecedorId int) ([]models.PedidoFornecedor, error)
}

type pedidoFornecedorService struct {
	repo PedidoFornecedorRepository
}

func NovoPedidoFornecedorService(repo PedidoFornecedorRepository) PedidoFornecedorService {
	return &pedidoFornecedorService{repo: repo}
}

func (s *pedidoFornecedorService) CriarPedidoFornecedor(pedido *models.PedidoFornecedor) error {
	return s.repo.CriarPedidoFornecedor(pedido)
}

func (s *pedidoFornecedorService) ListarPedidosPorFornecedor(fornecedorId int) ([]models.PedidoFornecedor, error) {
	return s.repo.ListarPedidosPorFornecedor(fornecedorId)
}
