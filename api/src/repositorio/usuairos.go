package repositorio

import (
	"api/src/model"
	"database/sql"
)

// representa um repositorio de usuários
type usuarios struct {
	db *sql.DB
}

// cria um repositorio de usuários
func NovoRepositorioUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// inserir um usuario no banco de dados
func (u usuarios) Criar(usuario model.Usuario) (uint64, error) {
	return 0, nil
}
