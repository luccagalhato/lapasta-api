package pedidofornecedor

import (
	database "lapasta/database"
	"lapasta/internal/models"
)

type PedidoFornecedorRepository interface {
	CriarPedidoFornecedor(pedido *models.PedidoFornecedor) error
	ListarPedidosPorFornecedor(fornecedorId int) ([]models.PedidoFornecedor, error)
}

type pedidoFornecedorRepository struct {
	db *database.SQLStr
}

func NovoPedidoFornecedorRepository(db *database.SQLStr) PedidoFornecedorRepository {
	return &pedidoFornecedorRepository{db: db}
}

func (r *pedidoFornecedorRepository) CriarPedidoFornecedor(pedido *models.PedidoFornecedor) error {
	return r.db.CriarPedidoFornecedor(pedido)
}

func (r *pedidoFornecedorRepository) ListarPedidosPorFornecedor(fornecedorId int) ([]models.PedidoFornecedor, error) {
	return r.db.ListarPedidosFornecedor(fornecedorId)
}
