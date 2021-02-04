package controllers

import (
	"devbook-api/src/autenticacao"
	"devbook-api/src/banco"
	"devbook-api/src/modelos"
	"devbook-api/src/repositorios"
	"devbook-api/src/respostas"
	"devbook-api/src/seguranca"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Login faz a funcao de autenticar o usuario para realizar login
func Login(w http.ResponseWriter, r *http.Request) {
	// lendo todo o corpo da requisicao para obter email e sneha

	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	// transformar o corpo da requisicao que eh um json em um usuario usando UNMARSHAL
	var usuario modelos.Usuario
	if erro = json.Unmarshal([]byte(corpoDaRequisicao), &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	fmt.Println(usuario)

	//abrir a conexao
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	//obtendo email e senha do usuario
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// obtendo token do login
	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)
	respostas.JSON(w, http.StatusOK, modelos.DadosAutenticacao{ID: usuarioID, Token: token})
}
