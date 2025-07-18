package documento

import (
	database "lapasta/database"
	"lapasta/internal/models"
	"time"
)

type DocumentoRepository interface {
	CriarDocumento(documento *models.Documento) error
	ListarDocumentos(page int) ([]models.Documento, error)
	FiltrarDataDocumento(inicioData, fimData time.Time) ([]models.Documento, error)
}

type documentoRepository struct {
	db *database.SQLStr
}

func NovoDocumentoRepository(db *database.SQLStr) DocumentoRepository {
	return &documentoRepository{
		db: db,
	}
}

func (r *documentoRepository) CriarDocumento(documento *models.Documento) error {
	return r.db.CriarDocumento(documento)
}

func (r *documentoRepository) ListarDocumentos(page int) ([]models.Documento, error) {
	return r.db.ListarDocumentos(page)
}

func (r *documentoRepository) FiltrarDataDocumento(inicioData, fimData time.Time) ([]models.Documento, error) {
	return r.db.FiltrarDataDocumento(inicioData, fimData)
}
