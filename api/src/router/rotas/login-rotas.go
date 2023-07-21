package rotas

import (
	"api/src/controller"
	"net/http"
)

var loginRota = Rota{
	URI:         "/login",
	Metodo:      http.MethodPost,
	Funcao:      controller.Login,
	Autenticado: false,
}
