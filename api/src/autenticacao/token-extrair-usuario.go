package autenticacao

// ExtrairUsuarioID retorna o usuarioId que está salvo no token
//func ExtrairUsuarioID(request *http.Request) (uint64, error) {
//	tokenString := extrairToken(request)
//	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
//	if erro != nil {
//		return 0, erro
//	}
//
//	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
//		if erro != nil {
//			return 0, erro
//		}
//
//		return usuarioID, nil
//	}
//
//	return 0, errors.New("Token inválido")
//}
