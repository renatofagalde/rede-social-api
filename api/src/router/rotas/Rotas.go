package rotas

import "net/http"

// Rota representa o modelo das rotas para as apis
type Rota struct {
	URI         string
	Metodo      string
	Funcao      func(http.ResponseWriter, *http.Request)
	Autenticado bool
}
