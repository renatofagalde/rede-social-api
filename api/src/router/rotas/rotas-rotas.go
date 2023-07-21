package rotas

import (
	"api/src/middleware"
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
	rotas := rotasUsuarios //definido no arquivo usuario.go dentro deste pacote
	rotas = append(rotas, loginRota)
	for _, rota := range rotas {
		if rota.Autenticado {
			r.HandleFunc(rota.URI,
				middleware.Logger(middleware.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI,
				middleware.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r
}
