package autenticacao

import (
	"api/src/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// criar token
func CriarToken(id uint64) (string, error) {

	permissoes := jwt.MapClaims{} //reivindicações
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["id"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes) //passando o método de aasinatura
	return token.SignedString(config.SecretKey)                    //secret
}
