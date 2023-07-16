package rotas

import (
	"api/src/controller"
	"net/http"
)

const rotaUsuarioId = "/usuarios/{id}"

var rotasUsuarios = []Rota{
	{
		URI:         "/usuarios",
		Metodo:      http.MethodPost,
		Funcao:      controller.Criar,
		Autenticado: false,
	},
	{
		URI:         "/usuarios",
		Metodo:      http.MethodGet,
		Funcao:      controller.BuscarTodos,
		Autenticado: true,
	},
	{
		URI:         "/usuarios/me",
		Metodo:      http.MethodGet,
		Funcao:      controller.BuscarLogado,
		Autenticado: true,
	},
	{
		URI:         rotaUsuarioId,
		Metodo:      http.MethodGet,
		Funcao:      controller.BuscarPorId,
		Autenticado: false,
	},
	{
		URI:         rotaUsuarioId,
		Metodo:      http.MethodPut,
		Funcao:      controller.Atualizar,
		Autenticado: false,
	},
	{
		URI:         rotaUsuarioId,
		Metodo:      http.MethodDelete,
		Funcao:      controller.Deletar,
		Autenticado: false,
	},
}
