package documento

import (
	"lapasta/internal/models"
	"time"
)

type DocumentoService interface {
	CriarDocumento(documento *models.Documento) error
	ListarDocumentos(page int) ([]models.Documento, error)
	FiltrarDataDocumento(inicioData, fimData time.Time) ([]models.Documento, error)
}

type documentoService struct {
	repo DocumentoRepository
}

func NovoDocumentoService(repo DocumentoRepository) DocumentoService {
	return &documentoService{
		repo: repo,
	}
}

func (s *documentoService) CriarDocumento(documento *models.Documento) error {
	return s.repo.CriarDocumento(documento)
}

func (s *documentoService) ListarDocumentos(page int) ([]models.Documento, error) {
	return s.repo.ListarDocumentos(page)
}
func (s *documentoService) FiltrarDataDocumento(inicioData, fimData time.Time) ([]models.Documento, error) {
	return s.repo.FiltrarDataDocumento(inicioData, fimData)
}
