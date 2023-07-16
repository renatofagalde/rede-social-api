package middleware

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"api/src/tool"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no console
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
		if len(request.Header["Trace-Id"]) == 0 {
			respostas.ERRO(writer, http.StatusBadRequest, errors.New("cabeçalho 'Trace-Id' é obrigatório"), http.StatusBadRequest)
			return
		}
		if erro := autenticacao.ValidarToken(request); erro != nil {
			fmt.Println(fmt.Sprintf("requisição NEGADA %v", request.Header["Trace-Id"]))
			respostas.ERRO(writer, http.StatusUnauthorized, erro, http.StatusUnauthorized)
			return
		}
		fmt.Println("requisição liberada")
		next(writer, request)
	}
}
