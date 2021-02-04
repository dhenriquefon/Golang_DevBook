package modelos

import "time"

//Publicacao representa uma publicacao feita por um usuario
type Publicacao struct {
	ID        uint64    `json: "id, omitempty"`
	Titulo    string    `json: "Titulo, omitempty`
	Conteudo  string    `json: "conteudo, omitempty`
	AutorID   uint64    `json: "autorId, omitempty`
	AutorNick string    `json: "autorNick, omitempty`
	Curtidas  uint64    `json: "curtidas, omitempty`
	CriadoEm  time.Time `json: "criadoEm, omitempty`
}
