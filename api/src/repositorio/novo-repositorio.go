package repositorio

import "database/sql"

// representa um repositorio de usuários
type usuarios struct {
	db *sql.DB
}

// cria um repositorio de usuários
func NovoRepositorioUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}
}
