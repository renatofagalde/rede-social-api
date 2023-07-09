package rotas

import "net/http"

const rotaUsuarioId = "/usuarios/{id}"

var rotasUsuarios = []Rota{
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: func(writer http.ResponseWriter, request *http.Request) {

		},
		Autenticado: false,
	},
	{
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: func(writer http.ResponseWriter, request *http.Request) {

		},
		Autenticado: false,
	},
	{
		URI:    rotaUsuarioId,
		Metodo: http.MethodGet,
		Funcao: func(writer http.ResponseWriter, request *http.Request) {

		},
		Autenticado: false,
	},
	{
		URI:    rotaUsuarioId,
		Metodo: http.MethodPut,
		Funcao: func(writer http.ResponseWriter, request *http.Request) {

		},
		Autenticado: false,
	},
	{
		URI:    rotaUsuarioId,
		Metodo: http.MethodDelete,
		Funcao: func(writer http.ResponseWriter, request *http.Request) {

		},
		Autenticado: false,
	},
}
