package controller

import "net/http"

// criar um usuario
func Atualizar(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("atualizando"))
}
