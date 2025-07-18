package boleto

import (
	database "lapasta/database"
	"lapasta/internal/models"
)

type BoletoRepository interface {
	CriarBoleto(boleto *models.Boleto) error
	ListarBoletosPorFornecedor(fornecedorId int) ([]models.Boleto, error)
	ListarBoletosPorPedido(pedidoId int) ([]models.Boleto, error)
	ListarBoletosDoDia(data string) ([]models.Boleto, error)
	PagarBoleto(codigoBarras string) error
	ListarBoletosVencidos() ([]models.Boleto, error)
	ListarBoletosPagos() ([]models.Boleto, error)
	ListarBoletosPendentes() ([]models.Boleto, error)
	AtualizarStatusBoleto(id int, statusId int) error
	GerarEEnviarRelatorioBoletos(emailAdmin string, ano, mes int, fornecedorID *int, statusIDs []int) error 
}

type boletoRepository struct {
	db *database.SQLStr
}

func NovoBoletoRepository(db *database.SQLStr) BoletoRepository {
	return &boletoRepository{
		db: db,
	}
}

func (r *boletoRepository) CriarBoleto(boleto *models.Boleto) error {
	return r.db.CriarBoletoRecebido(boleto)
}

func (r *boletoRepository) ListarBoletosPorFornecedor(fornecedorId int) ([]models.Boleto, error) {
	return r.db.ListarBoletosPorFornecedor(fornecedorId)
}

func (r *boletoRepository) ListarBoletosPorPedido(pedidoId int) ([]models.Boleto, error) {
	return r.db.ListarBoletosPorPedido(pedidoId)
}

func (r *boletoRepository) ListarBoletosDoDia(data string) ([]models.Boleto, error) {
	return r.db.ListarBoletosDoDia(data)
}

func (r *boletoRepository) PagarBoleto(codigoBarras string) error {
	return r.db.PagarBoleto(codigoBarras)
}

func (r *boletoRepository) ListarBoletosPagos() ([]models.Boleto, error) {
	return r.db.ListarBoletosPagos()
}

func (r *boletoRepository) ListarBoletosVencidos() ([]models.Boleto, error) {
	return r.db.ListarBoletosVencidos()
}
func (r *boletoRepository) ListarBoletosPendentes() ([]models.Boleto, error) {
	return r.db.ListarBoletosPendentes()
}
func (r *boletoRepository) AtualizarStatusBoleto(id int, statusId int) error {
	return r.db.AtualizarStatusBoleto(id, statusId)
}
func (r *boletoRepository) GerarEEnviarRelatorioBoletos(emailAdmin string, ano, mes int, fornecedorID *int, statusIDs []int) error  {
	return r.db.GerarEEnviarRelatorioBoletos(emailAdmin, ano, mes, fornecedorID, statusIDs)
}
