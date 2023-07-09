package controller

import "net/http"

func BuscarTodos(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("criando usuario"))
}
