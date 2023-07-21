package controller

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/repositorio"
	_ "api/src/repositorio"
	"api/src/respostas"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DeixarDeSeguir(write http.ResponseWriter, request *http.Request) {

	//usuario do token
	seguidorID, erro := autenticacao.ExtrairUsuarioID(request)
	if erro != nil {
		respostas.ERRO(write, http.StatusUnauthorized, erro, http.StatusUnauthorized)
		return
	}

	parametros := mux.Vars(request)

	usuarioID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		respostas.ERRO(write, http.StatusBadRequest, erro, http.StatusBadRequest)
		return
	}

	if seguidorID == usuarioID {
		respostas.ERRO(write, http.StatusForbidden, errors.New("Você não pode seguir você mesmo"), http.StatusForbidden)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro, http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuario(db)
	if erro = repositorio.DeixarDeSeguir(usuarioID, seguidorID); erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro, http.StatusInternalServerError)
		return
	}

	respostas.JSON(write, http.StatusNoContent, nil)
}
