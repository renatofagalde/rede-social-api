package controller

import (
	"api/src/banco"
	"api/src/repositorio"
	"api/src/respostas"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// buscando um usuario por id
func BuscarPorId(write http.ResponseWriter, request *http.Request) {
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
	usuario, erro := respositorio.BuscarPorID(ID)
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro, http.StatusInternalServerError)
	}
	respostas.JSON(write, http.StatusOK, usuario)
}
