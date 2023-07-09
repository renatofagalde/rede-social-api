package controller

import "net/http"

// buscando um usuario por id
func BuscarPorId(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("criando usuario"))
}
