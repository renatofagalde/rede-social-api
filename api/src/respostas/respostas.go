package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON -> retora um resposta em JSON
func JSON(writer http.ResponseWriter, statusCode int, dados interface{}) {
	writer.WriteHeader(statusCode)
	writer.Header().Add("versao", "01")
	writer.Header().Add("Content-Type", "application/json")
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
