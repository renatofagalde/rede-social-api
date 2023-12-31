package model

import (
	"api/src/seguranca"
	"errors"
	"fmt"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

// Usuario modelo
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

// validar propriedades se estão preenchidas
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório")
	}

	if usuario.Email == "" {
		return errors.New("Email é obrigatório")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New(fmt.Sprintf("O email '%s' é inválido", usuario.Email))
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("Senha é obrigatório")
	}
	return nil
}

// remover espacos das propriedades
func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}

// ira chamar os metodos de formatar e checagem
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}
