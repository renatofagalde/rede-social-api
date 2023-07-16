package controller

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/model"
	"api/src/repositorio"
	_ "api/src/repositorio"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io"
	"net/http"
)

// criar login para um  usuario
func Login(write http.ResponseWriter, request *http.Request) {
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

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuario(db)
	usuarioSalvo, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro)
		return
	}

	//valida se a senha possui o mesmo valor
	if erro = seguranca.VerificarSenha(usuarioSalvo.Senha, usuario.Senha); erro != nil {
		respostas.ERRO(write, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuario.ID)
	if erro != nil {
		respostas.ERRO(write, http.StatusInternalServerError, erro)
		return
	}
	write.Write([]byte(token))
}
