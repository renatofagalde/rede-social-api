package repositorio

import (
	"api/src/model"
	"fmt"
)

// trazendo usuarios que atendem o filtro por nome ou nick
func (u usuarios) BuscarPorNomeOuNick(nomeOuNick string) ([]model.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //escape de caracteres

	cursor, erro := u.db.Query("select id, nome, nick, email, criado_em from usuarios "+
		"where nome like ? or nick like ?",
		nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer cursor.Close()

	var usuarios []model.Usuario
	for cursor.Next() {
		var usuario model.Usuario
		if erro := cursor.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
