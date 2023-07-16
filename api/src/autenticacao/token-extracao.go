package autenticacao

import (
	"net/http"
	"strings"
)

func extrairToken(request *http.Request) string {
	token := request.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}
