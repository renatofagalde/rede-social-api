package controller

import "net/http"

func Criar(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("criando"))
}
