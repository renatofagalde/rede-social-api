package controller

import "net/http"

func Deletar(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("deletando"))
}
