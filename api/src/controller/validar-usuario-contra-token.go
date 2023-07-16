package controller

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"errors"
	"net/http"
)

func validarUsuario(write http.ResponseWriter, request *http.Request, erro error, ID uint64) bool {
	usuarioIdToken, erro := autenticacao.ExtrairUsuarioID(request)
	if erro != nil {
		respostas.ERRO(write, http.StatusUnauthorized, erro, http.StatusUnauthorized)
		return true
	}

	if ID != usuarioIdToken {
		respostas.ERRO(write, http.StatusForbidden, errors.New("sem permissão"), http.StatusForbidden)
		return true
	}
	return false
}
