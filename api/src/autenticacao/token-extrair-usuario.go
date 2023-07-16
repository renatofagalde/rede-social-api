package autenticacao

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
)

// ExtrairUsuarioID retorna o usuarioId que está salvo no token
func ExtrairUsuarioID(request *http.Request) (uint64, error) {
	tokenString := extrairToken(request)
	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["id"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return id, nil
	}
	return 0, errors.New("token inválido")
}
