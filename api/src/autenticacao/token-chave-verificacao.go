package autenticacao

import (
	"api/src/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// validar a se o tipo do toke pertence a familia de assinaturas existentes
// isso está na documentação doGO
func retornarChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		fmt.Sprintf("Erro ao validar algoritmo: %s", token.Header["alg"])
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
