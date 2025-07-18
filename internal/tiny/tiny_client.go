package tiny

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	sql "lapasta/database"
	"lapasta/internal/models"
	"net/http"
)

type TinyClient struct {
	Token string
}

func NewTinyClient(sqlConn *sql.SQLStr) (*TinyClient, error) {
	token, err := sqlConn.GetTokenTiny()
	if err != nil {
		return nil, err
	}
	return &TinyClient{Token: token}, nil
}

func (c *TinyClient) BuscarExpedicoesPorAgrupamento(idAgrupamento string) (*models.TinyExpedicaoResponse, error) {
	url := fmt.Sprintf("https://api.tiny.com.br/api2/expedicao.pesquisar.agrupamentos.php?token=%s&idAgrupamento=%s", c.Token, idAgrupamento)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("erro na resposta da API do Tiny")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var resultado models.TinyExpedicaoResponse
	err = json.Unmarshal(body, &resultado)
	if err != nil {
		return nil, err
	}

	return &resultado, nil
}
