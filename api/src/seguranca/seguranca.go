package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// valida se uma senha em branco tem o mesmo valor que uma senha cryptografada
func VerificarSenha(hash string, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
}
