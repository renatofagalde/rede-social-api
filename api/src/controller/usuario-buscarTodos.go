package controller

import "net/http"

func buscarTodos(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("criando usuario"))
}
