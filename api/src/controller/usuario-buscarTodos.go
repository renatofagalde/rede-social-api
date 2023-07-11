package controller

import (
	"api/src/banco"
	"api/src/repositorio"
	"api/src/respostas"
	"net/http"
	"strings"
)

func BuscarTodos(write http.ResponseWriter, request *http.Request) {
	nomeOuNick := strings.ToLower(request.URL.Query().Get("usuario"))
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro)
	}
	defer db.Close()
	respositorio := repositorio.NovoRepositorioUsuario(db)
	usuarios, erro := respositorio.BuscarPorNomeOuNick(nomeOuNick)
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro)
	}

	respostas.JSON(write, http.StatusOK, usuarios)

}
