package middleware

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"api/src/tool"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var headers string

		for k, v := range request.Header {
			h := fmt.Sprintf("%v: %v\n", k, v)
			headers += h
		}

		payload := io.NopCloser(tool.ReusableReader(request.Body))
		request.Body = payload

		payloadLog, _ := io.ReadAll(payload)

		log.Printf("\n\n********* INICIO\nhost: %s, "+
			"URI: %s, "+
			"method: %s,"+
			"\nheaders:\n%s=========\n"+
			"payloadLog:\n"+
			"%s\n********* FIM",
			request.Host, request.RequestURI, request.Method, headers, payloadLog)

		next(writer, request)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if erro := autenticacao.ValidarToken(request); erro != nil {
			respostas.ERRO(writer, http.StatusUnauthorized, erro)
		}
		fmt.Println("autenticanção liberado")
		next(writer, request)
	}
}
