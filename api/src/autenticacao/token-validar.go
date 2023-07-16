package autenticacao

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

// verifica se o token passado na requisição é válido
func ValidarToken(request *http.Request) error {
	tokenString := extrairToken(request)
	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil {
		return erro
	}
	fmt.Println(token)
	return nil
}
