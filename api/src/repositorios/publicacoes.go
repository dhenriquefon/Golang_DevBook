package repositorios

import (
	"database/sql"
	"devbook-api/src/modelos"
	"fmt"
)

type Publicacoes struct {
	db *sql.DB
}

//NovoRepositorioDePublicacoes cria um repositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicacao no banco de dados
func (repoPublicacao Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repoPublicacao.db.Prepare(
		"INSERT into publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoID, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoID), nil

}

//Buscar retorna as publicacoes do usuario, e tambem dos usuarios que ele segue
func (repoPublicacao Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repoPublicacao.db.Query(`
SELECT DISTINCT p.*, u.nick FROM publicacoes p
JOIN usuarios u ON u.id = p.autor_id JOIN seguidores s
ON p.autor_id = s.usuario_id WHERE u.id = ? OR s.seguidor_id = ? ORDER BY 1 desc
		`, usuarioID, usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil

}

//BuscarPorID retorna a publicacao de um ID
func (repoPublicacao Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linhas, erro := repoPublicacao.db.Query(
		"SELECT p.*, u.nick FROM publicacoes p JOIN usuarios u ON u.id = p.autor_id WHERE p.id = ?",
		publicacaoID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	var publicacao modelos.Publicacao

	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil

}

// Atualizar atualiza uma publicacao no banco de dados
func (repoPublicacao Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {

	fmt.Println(publicacaoID)
	fmt.Println(publicacao)
	statement, erro := repoPublicacao.db.Prepare(`
		UPDATE publicacoes SET
			titulo = ?,
			conteudo = ?
		WHERE
			ID = ?
		`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID)
	if erro != nil {
		return erro
	}

	return nil

}

// Excluir exclui uma publicacao
func (repoPublicacao Publicacoes) Excluir(publicacaoID uint64) error {

	statement, erro := repoPublicacao.db.Prepare(`
		DELETE FROM publicacoes WHERE ID = ?
		`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil

}

//BuscarPorUsuario retorna a publicacao de um ID
func (repoPublicacao Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repoPublicacao.db.Query(
		"SELECT p.*, u.nick FROM publicacoes p JOIN usuarios u ON u.id = p.autor_id WHERE p.autor_id = ?",
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil

}

// Curtir curte uma publicacao
func (repoPublicacao Publicacoes) Curtir(publicacaoID uint64) error {

	statement, erro := repoPublicacao.db.Prepare(`
		UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?
		`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil

}

// Descurtir descurte uma publicacao
func (repoPublicacao Publicacoes) Descurtir(publicacaoID uint64) error {

	statement, erro := repoPublicacao.db.Prepare(`
		UPDATE publicacoes SET curtidas = CASE WHEN curtidas > 0 THEN curtidas - 1 ELSE 0 END WHERE id = ?
		`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil

}
