package controller

import "net/http"

func criar(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("criando usuario"))
}
