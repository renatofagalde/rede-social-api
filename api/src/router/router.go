package router

import (
	"api/src/router/rotas"
	"github.com/gorilla/mux"
)

// Configurações das rotas
func Gerar() *mux.Router {

	r := mux.NewRouter()
	return rotas.Configurar(r)
}
