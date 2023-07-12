package repositorio

import (
	"api/src/model"
)

// traz um usuário do banco de dados usuando o email com filtro
func (u usuarios) BuscarPorEmail(email string) (model.Usuario, error) {

	cursor, erro := u.db.Query("select id,senha from usuarios "+
		"where email =?", email)
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer cursor.Close()

	var usuario model.Usuario
	for cursor.Next() {
		if erro := cursor.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return model.Usuario{}, erro
		}
	}
	return usuario, nil
}
