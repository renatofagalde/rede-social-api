package controller

import "net/http"

func deletar(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("criando usuario"))
}
