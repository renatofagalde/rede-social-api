package middleware

import (
	"fmt"
	"log"
	"net/http"
)

// Logger escreve logs
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("\n********* INICIo\nMethod: %s, URI: %s, HOST: %s\n*********",
			request.Method, request.RequestURI, request.Host)
	}
}

// autenticar
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("validando o token")
		next(writer, request) //next -> chamar a execução que veio no parametro que veio do if da rota quando autenticado for true
	}
}
