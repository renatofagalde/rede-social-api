package controller

import (
	"api/src/banco"
	"api/src/model"
	"api/src/repositorio"
	_ "api/src/repositorio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Criar(writer http.ResponseWriter, request *http.Request) {
	payload, erro := io.ReadAll(request.Body)
	if erro != nil {
		log.Fatal(erro)
		return
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(payload, &usuario); erro != nil {
		writer.Write([]byte("Erro ao converter usuário para struct"))
	}
	fmt.Println(usuario)

	db, erro := banco.Conectar()
	if erro != nil {
		writer.Write([]byte("Erro ao conectar com o BD"))
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuario(db)
	repositorio.Criar(usuario)

}
