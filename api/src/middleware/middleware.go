package middleware

import (
	"api/src/respostas"
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

		payload, erro := io.ReadAll(request.Body)
		if erro != nil {
			respostas.ERRO(writer, http.StatusUnprocessableEntity, erro)
			return
		}
		log.Printf("\n\n********* INICIO\nmethod: %s, URI: %s, host: %s,\nheaders:\n%s=========\npayload:\n%s\n********* FIM",
			request.Method, request.RequestURI, request.Host, headers, payload)

		next(writer, request)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("autenticando")
		next(writer, request)
	}
}
