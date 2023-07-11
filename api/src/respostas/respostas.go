package respostas

import (
	"api/src/config"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// JSON -> retora um resposta em JSON
func JSON(writer http.ResponseWriter, statusCode int, dados interface{}) {
	writer.Header().Set("versao", strconv.Itoa(config.VersaoAPI))
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	if erro := json.NewEncoder(writer).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// retorna um erro em JSON
func ERRO(writer http.ResponseWriter, statusCode int, erro error) {
	JSON(writer, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
