package repositorio

// Deletar um usuario
func (u usuarios) DeletarPorId(ID uint64) error {
	statement, erro := u.db.Prepare("delete from usuarios where id=?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}
	return nil
}
