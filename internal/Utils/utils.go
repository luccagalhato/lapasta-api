package utils

import (
	_ "embed" // EMB
	sql "lapasta/database"
	"log"
	"math/rand"
	"time"
)

// ConnectionDb ..
var ConnectionDb *sql.SQLStr

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// SetSQLConn armazena a conexão com o banco de dados.
func SetSQLConn(l *sql.SQLStr) {
	if l == nil {
		log.Fatal("A conexão com o banco de dados não pode ser nula.")
	}
	ConnectionDb = l
}

// GerarStringAleatoria gera uma string aleatória de comprimento fixo
func GerarStringAleatoria(tamanho int) string {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	b := make([]byte, tamanho)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}
