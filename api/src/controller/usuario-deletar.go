package controller

import (
	"api/src/banco"
	"api/src/repositorio"
	"api/src/respostas"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Deletar(write http.ResponseWriter, request *http.Request) {
	parametros := mux.Vars(request)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		respostas.ERRO(write, http.StatusBadRequest, erro, http.StatusBadRequest)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro, http.StatusInternalServerError)
	}
	defer db.Close()

	respositorio := repositorio.NovoRepositorioUsuario(db)
	if erro = respositorio.DeletarPorId(ID); erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro, http.StatusInternalServerError)
	}
	respostas.JSON(write, http.StatusNoContent, nil)
}
