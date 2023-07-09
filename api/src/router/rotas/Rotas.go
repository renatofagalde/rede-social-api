package rotas

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Rota representa o modelo das rotas para as apis
type Rota struct {
	URI         string
	Metodo      string
	Funcao      func(http.ResponseWriter, *http.Request)
	Autenticado bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
