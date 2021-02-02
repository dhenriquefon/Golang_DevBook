package modelos

import (
	"devbook-api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//Usuario
type Usuario struct {
	ID       uint64    `json: "id, omitempty"` // omitempty : caso o campo nao exista no json nao vai ser inserido 0
	Nome     string    `json: "nome, omitempty"`
	Nick     string    `json: "nick, omitempty"`
	Email    string    `json: "email, omitempty"`
	Senha    string    `json: "senha, omitempty"`
	CriadoEm time.Time `json: "CriadoEm, omitempty"`
}

// Montando as validacoes

//Preparar vai chamar os metodos para validar e formatar o usuario
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email nao eta em uma formato valido")
	}

	if etapa == "cadastro" {
		if usuario.Senha == "" {
			return errors.New("A senha é obrigatória e não pode estar em branco")
		}
	}

	return nil
}

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
