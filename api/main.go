package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()
	fmt.Println(fmt.Sprintf("URL banco de dados = %s", config.StringURLBanco))
	fmt.Println(fmt.Sprintf("Porta da aplicacao = %d", config.Porta))

	fmt.Println("Rodando api")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
