package repositorio

import (
	"api/src/model"
)

// traz um usuário do banco de dados
func (u usuarios) BuscarPorID(ID uint64) (model.Usuario, error) {

	cursor, erro := u.db.Query("select id, nome, nick, email, criado_em from usuarios "+
		"where id =?", ID)
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer cursor.Close()

	var usuario model.Usuario
	for cursor.Next() {
		if erro := cursor.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return model.Usuario{}, erro
		}
	}
	return usuario, nil
}
