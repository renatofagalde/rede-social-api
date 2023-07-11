package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	// URL de conexao com o banco de dados
	StringURLBanco = ""
	//Porta que irá subir a aplicação
	Porta = 0
)

// Carregar irá inicializar as variáveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	// os.Getenv() vem do pacote godotenv
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	//"aplicacao:1212@tcp(localhost:3300)/dev?charset=utf8&parseTime=true&loc=Local"
	StringURLBanco = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))
}
