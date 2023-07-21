package repositorio

// deixar de seguir
func (u usuarios) DeixarDeSeguir(usuarioID uint64, seguirID uint64) error {

	statement, erro := u.db.Prepare("delete from seguidores where usuario_id=? and seguidor_id=?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuarioID, seguirID); erro != nil {
		return erro
	}
	return nil
}
