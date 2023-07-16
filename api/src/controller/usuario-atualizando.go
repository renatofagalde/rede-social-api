package controller

import (
	"api/src/banco"
	"api/src/model"
	"api/src/repositorio"
	"api/src/respostas"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

// criar um usuario
func Atualizar(write http.ResponseWriter, request *http.Request) {
	parametros := mux.Vars(request)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		respostas.ERRO(write, http.StatusBadRequest, erro, http.StatusBadRequest)
		return
	}

	payload, erro := io.ReadAll(request.Body)
	if erro != nil {
		respostas.ERRO(write, http.StatusUnprocessableEntity, erro, http.StatusUnprocessableEntity)
		return
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(payload, &usuario); erro != nil {
		respostas.ERRO(write, http.StatusBadRequest, erro, http.StatusBadRequest)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.ERRO(write, http.StatusBadRequest, erro, http.StatusBadRequest)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro, http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuario(db)
	if erro = repositorio.Atualizar(ID, usuario); erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro, http.StatusInternalServerError)
		return
	}
	respostas.JSON(write, http.StatusNoContent, nil)
}
