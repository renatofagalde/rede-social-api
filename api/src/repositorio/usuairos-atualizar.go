package repositorio

import (
	"api/src/model"
)

// Atualizar um usuario
func (u usuarios) Atualizar(ID uint64, usuario model.Usuario) error {
	statement, erro := u.db.Prepare("update usuarios set nome=?, nick=?, email=? where id=?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}
	return nil
}
