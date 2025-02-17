package modelos

import (
	"errors"
	"strings"
	"time"
)

// Representa uma publicação feita por um usuario
type Publicacao struct {
	ID        uint64    `json: "id, omitempty"`
	Titulo    string    `json: "titulo, omitempty"`
	Conteudo  string    `json: "conteudo, omitempty"`
	AutorID   uint64    `json: "autorId, omitempty"`
	AutorNick uint64    `json: "autorNick, omitempty"`
	Curtidas  uint64    `json: "curtidas"`
	CriadaEm  time.Time `json: "criadaEm, omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar a publicacao recebida
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o titulo é obrigatorio e não pode estar em branco")
	}

	if publicacao.Conteudo == "" {
		return errors.New("o Conteudo é obrigatorio e não pode estar em branco")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
}
