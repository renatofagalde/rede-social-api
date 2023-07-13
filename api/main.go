package main

import (
	"api/src/config"
	"api/src/router"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

/*func init() {
	chave := make([]byte, 64)
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}
*/

func main() {

	config.Carregar()
	fmt.Println(fmt.Sprintf("URL banco de dados = %s", config.StringURLBanco))
	fmt.Println(fmt.Sprintf("Porta da aplicacao = %d", config.Porta))

	fmt.Println("Rodando api")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
