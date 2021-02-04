package repositorios

// repositorios eh quem vai de fato interagir com banco de dados, queries, etc
import (
	"database/sql"
	"devbook-api/src/modelos"
	"fmt"
)

// Usuarios representa um repositorio de Usuarios
type Usuarios struct {
	db *sql.DB
}

//NovoRepositorioDeUsuarios cria uma instancia com o banco que foi aberto
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Implementando as funcoes do meu repositorio usuario

//Criar (detalhe primeiro parametro repoUsuarios Usuarios indica que minha funcao esta dentro do repositorio/tipo Usuarios) . NOme da funcao eh Criar e retorna uint64, error
func (repoUsuarios Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repoUsuarios.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	//executar o statement
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDinserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDinserido), nil
}

// Buscar traz todos os usuarios com filtro de nome ou nick
func (repoUsuarios Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repoUsuarios.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

//BuscarPorID busca usuario por ID
func (repoUsuarios Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repoUsuarios.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id= ?", ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar atualiza os dados de um usuario
func (repoUsuarios Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repoUsuarios.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		fmt.Println(erro)
		return erro
	}

	return nil
}

// Excluir exclui um usuarios da base de dados
func (repoUsuarios Usuarios) Excluir(ID uint64) error {
	statement, erro := repoUsuarios.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca o usuario pelo email
func (repoUsuarios Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linhas, erro := repoUsuarios.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		erro = linhas.Scan(&usuario.ID, &usuario.Senha)
		if erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

//SeguirUsuario permite que um usuario siga o outro inserindo o dado na tabela seguidores
func (repoUsuarios Usuarios) SeguirUsuario(usuarioID uint64, seguidorID uint64) error {
	statement, erro := repoUsuarios.db.Prepare(
		"INSERT ignore INTO seguidores (usuario_id, seguidor_id) VALUES (?, ?)",
	)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuarioID, seguidorID)
	if erro != nil {
		return erro
	}

	return nil
}

//PararDeSeguirUsuario permite que um usuario siga o outro inserindo o dado na tabela seguidores
func (repoUsuarios Usuarios) PararDeSeguirUsuario(usuarioID uint64, seguidorID uint64) error {
	statement, erro := repoUsuarios.db.Prepare(
		"DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?",
	)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuarioID, seguidorID)
	if erro != nil {
		return erro
	}

	return nil
}

//BuscarSeguidores retorna todos os seguidores de um usuario
func (repoUsuarios Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, erro := repoUsuarios.db.Query(`
SELECT u.ID, u.nome, u.nick, u.email , u.criadoEm FROM usuarios u
JOIN
seguidores s
on u.ID = s.seguidor_id
WHERE
s.usuario_id = ?
		`, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

//BuscarSeguindo retorna todos os usuarios quem seguem um usuario
func (repoUsuarios Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, erro := repoUsuarios.db.Query(`
SELECT u.ID, u.nome, u.nick, u.email , u.criadoEm FROM usuarios u
JOIN
seguidores s
on u.ID = s.usuario_id
WHERE
s.seguidor_id = ?
		`, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

//BuscarSenha busca a senha pelo ID pelo usuario
func (repoUsuarios Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repoUsuarios.db.Query("SELECT senha FROM usuarios WHERE ID=?", usuarioID)
	if erro != nil {
		return "", erro
	}

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, erro
}

//AtualizarSenha atualizar a senha do usuario
func (repoUsuarios Usuarios) AtualizarSenha(ID uint64, senhaComHash string) error {
	fmt.Println(senhaComHash)

	statement, erro := repoUsuarios.db.Prepare("UPDATE usuarios SET senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senhaComHash, ID); erro != nil {
		return erro
	}

	return nil
}
