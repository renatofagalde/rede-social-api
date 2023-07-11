package controller

import (
	"api/src/banco"
	"api/src/model"
	"api/src/repositorio"
	_ "api/src/repositorio"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Criar(write http.ResponseWriter, request *http.Request) {
	payload, erro := io.ReadAll(request.Body)
	if erro != nil {
		respostas.ERRO(write, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(payload, &usuario); erro != nil {
		respostas.ERRO(write, http.StatusBadRequest, erro)
		return

	}
	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.ERRO(write, http.StatusBadRequest, erro)
		return
	}
	fmt.Println(usuario)

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuario(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(write, http.StatusCreated, usuario)
}
