package repositorio

// inserir um seguidor
func (u usuarios) Seguir(usuarioID uint64, seguirID uint64) error {

	statement, erro := u.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuarioID, seguirID); erro != nil {
		return erro
	}
	return nil
}
