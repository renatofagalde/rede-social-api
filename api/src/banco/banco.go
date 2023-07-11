package banco

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // import implicito
	"log"
)

// Conexao com o banco de dados
// neste retorno apenas um pode existir, pq eles são mutualmente exclusivos.
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringURLBanco)
	if erro != nil {
		log.Printf("Erro ao conectar: %s", erro)
		return nil, erro //retorno desta função é exclusivamente
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
