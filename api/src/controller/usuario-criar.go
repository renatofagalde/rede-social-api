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
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(payload, &usuario); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(usuario)

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuario(db)
	ID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	writer.Write([]byte(fmt.Sprintf("ID inserido é: %d", ID)))

}
