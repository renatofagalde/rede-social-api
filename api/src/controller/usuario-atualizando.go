package controller

import "net/http"

// criar um usuario
func atualizar(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("criando usuario"))
}
