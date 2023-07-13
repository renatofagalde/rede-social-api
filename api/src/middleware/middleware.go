package middleware

import (
	"api/src/respostas"
	"io"
	"log"
	"net/http"
)

// Logger escreve logs
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		payload, erro := io.ReadAll(request.Body)
		if erro != nil {
			respostas.ERRO(writer, http.StatusUnprocessableEntity, erro)
			return
		}
		log.Printf("\n\n********* INICIO\nMethod: %s, URI: %s, host: %s, payload:\n%s\n********* FIM",
			request.Method, request.RequestURI, request.Host, payload)
	}
}

// autenticar um usuario
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		log.Printf("\n\n********* mamae")
		next(writer, request)
	}
}
