package repositorio

import (
	"api/src/model"
)

// inserir um usuario no banco de dados
func (u usuarios) Criar(usuario model.Usuario) (uint64, error) {
	statement, erro := u.db.Prepare("insert into usuarios (nome, nick, email, senha) value (?,?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}
	ID, erro := insercao.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ID), nil
}
