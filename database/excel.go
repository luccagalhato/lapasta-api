package sql

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"lapasta/internal/models"
	"mime/multipart"
	"mime/quotedprintable"
	"net/smtp"
	"net/textproto"
	"time"
)

func (s *SQLStr) GerarRelatorioMensal(mes int, ano int, emailAdmin string) (string, error) {
	if mes == 0 {
		mes = int(time.Now().Month())
		ano = time.Now().Year()
	}

	query := `
		SELECT 
			p.IdFuncionario, f.Nome, p.Dia, p.HManha, p.HAlmocoSaida, p.HAlmocoRetorno, p.HNoite
		FROM Ponto p
		JOIN Funcionarios f ON p.IdFuncionario = f.Id
		WHERE MONTH(p.Dia) = @Mes AND YEAR(p.Dia) = @Ano
		ORDER BY p.IdFuncionario, p.Dia
	`

	rows, err := s.db.Query(query, sql.Named("Mes", mes), sql.Named("Ano", ano))
	if err != nil {
		return "", fmt.Errorf("erro ao consultar pontos do mês: %w", err)
	}
	defer rows.Close()

	relatorio := make(map[int][]models.Ponto)
	for rows.Next() {
		var ponto models.Ponto
		if err := rows.Scan(&ponto.IdFuncionario, &ponto.NomeFuncionario, &ponto.Dia, &ponto.HManha, &ponto.HAlmocoSaida, &ponto.HAlmocoRetorno, &ponto.HNoite); err != nil {
			return "", fmt.Errorf("erro ao escanear ponto: %w", err)
		}
		relatorio[ponto.IdFuncionario] = append(relatorio[ponto.IdFuncionario], ponto)
	}
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	writer.Comma = ';'

	header := []string{"IdFuncionario", "Nome", "Dia", "HManha", "HAlmocoSaida", "HAlmocoRetorno", "HNoite"}
	if err := writer.Write(header); err != nil {
		return "", fmt.Errorf("erro ao escrever cabeçalho: %w", err)
	}

	formatarHorario := func(hora sql.NullString) string {
		if !hora.Valid || hora.String == "0001-01-01T00:00:00Z" || hora.String == "" {
			return " - "
		}
		t, err := time.Parse("0001-01-01T15:04:05Z", hora.String)
		if err != nil {
			return " - "
		}
		return t.Format("15:04:05")
	}

	var ultimoIdFuncionario int
	for _, pontos := range relatorio {
		for i, ponto := range pontos {
			if i == 0 && ultimoIdFuncionario != 0 {
				if err := writer.Write([]string{}); err != nil {
					return "", fmt.Errorf("erro ao escrever linha em branco no buffer: %w", err)
				}
			}

			var dataFormatada string
			if t, err := time.Parse("2006-01-02T15:04:05Z", ponto.Dia); err == nil {
				dataFormatada = t.Format("02/01/2006")
			} else {
				dataFormatada = "Null"
			}

			linha := []string{
				fmt.Sprintf("%d", ponto.IdFuncionario),
				ponto.NomeFuncionario,
				dataFormatada,
				formatarHorario(ponto.HManha),
				formatarHorario(ponto.HAlmocoSaida),
				formatarHorario(ponto.HAlmocoRetorno),
				formatarHorario(ponto.HNoite),
			}

			if err := writer.Write(linha); err != nil {
				return "", fmt.Errorf("erro ao escrever linha no buffer: %w", err)
			}

			ultimoIdFuncionario = ponto.IdFuncionario
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", fmt.Errorf("erro ao finalizar escrita no buffer: %w", err)
	}

	if err := s.enviarEmailComAnexo(emailAdmin, buf.Bytes(), mes, ano); err != nil {
		return "", fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	return "Relatório enviado com sucesso!", nil
}

func (s *SQLStr) enviarEmailComAnexo(emailAdmin string, fileContent []byte, mes, ano int) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	senderEmail := "email@gmail.com"
	senderPassword := "senha"
	subject := fmt.Sprintf("Relatório Mensal de Pontos - %02d/%d", mes, ano)
	body := fmt.Sprintf("Olá, em anexo está o relatório mensal de pontos do mês %02d do ano %d.", mes, ano)

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	textPart, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Type":              []string{"text/plain; charset=UTF-8"},
		"Content-Transfer-Encoding": []string{"quoted-printable"},
	})
	if err != nil {
		return fmt.Errorf("erro ao criar a parte do corpo do e-mail: %w", err)
	}
	encoder := quotedprintable.NewWriter(textPart)
	_, err = encoder.Write([]byte(body))
	if err != nil {
		return fmt.Errorf("erro ao escrever o corpo do e-mail: %w", err)
	}
	encoder.Close()

	filePart, err := writer.CreateFormFile("attachment", "relatorio_pontos.csv")
	if err != nil {
		return fmt.Errorf("erro ao criar a parte do anexo: %w", err)
	}
	_, err = filePart.Write(fileContent)
	if err != nil {
		return fmt.Errorf("erro ao escrever o anexo: %w", err)
	}
	writer.Close()

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	to := []string{emailAdmin}	
	msg := []byte("From: " + senderEmail + "\r\n" +
		"To: " + emailAdmin + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: multipart/mixed; boundary=\"" + writer.Boundary() + "\"\r\n" +
		"\r\n" + buf.String())

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, to, msg)
	if err != nil {
		return fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	return nil
}
