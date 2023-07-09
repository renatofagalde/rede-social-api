package router

import "github.com/gorilla/mux"

// Configurações das rotas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
