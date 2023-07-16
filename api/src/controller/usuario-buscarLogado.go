package controller

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/repositorio"
	"api/src/respostas"
	"net/http"
)

// buscando um usuario por id
func BuscarLogado(write http.ResponseWriter, request *http.Request) {

	ID, erro := autenticacao.ExtrairUsuarioID(request)
	if erro != nil {
		respostas.ERRO(write, http.StatusUnauthorized, erro, http.StatusUnauthorized)
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro, http.StatusInternalServerError)
	}
	defer db.Close()

	respositorio := repositorio.NovoRepositorioUsuario(db)
	usuario, erro := respositorio.BuscarPorID(ID)
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro, http.StatusInternalServerError)
	}
	respostas.JSON(write, http.StatusOK, usuario)
}
