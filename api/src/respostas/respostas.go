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
	writer.Header().Set("Versao", strconv.Itoa(config.VersaoAPI))
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	//tratar quando o tipo da resposta é 204
	if dados != nil {
		if erro := json.NewEncoder(writer).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}

}

// retorna um erro em JSON
func ERRO(writer http.ResponseWriter, statusCode int, erro error, codigo int) {
	JSON(writer, statusCode, struct {
		Erro   string `json:"erro"`
		Codigo int    `json:"codigo"`
	}{
		Erro:   erro.Error(),
		Codigo: codigo,
	})
}
