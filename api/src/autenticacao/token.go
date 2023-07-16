package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(id uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 8).Unix()
	permissoes["id"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}
