package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// CriarUsuario chama a API para cadastrar o usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// vamos fazer o mesmo papel do POSTMAN, criar uma requisicao e enviar
	// aqui estamos pegando os campos do formulario
	r.ParseForm()

	// Podemos obter os dados assim, mas mais produtivo eh transformar os dados em um MAPA para aplicar um Marshal para converter a requisicao
	// em JSON para chamar para a API
	//nome := r.FormValue("nome")
	//email := r.FormValue("email")
	//nick := r.FormValue("nick")
	//senha := r.FormValue("senha")

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	// imprimindo usuario em formato JSON
	//fmt.Println(bytes.NewBuffer(usuario))

	// fazendo a chamada para API
	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	//SEMPRE RODAR
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

// PararDeSeguirUsuario para de seguir um usuario
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	fmt.Println("teste douglas")

	// fazendo a chamada para API
	url := fmt.Sprintf("%s/usuarios/%d/parar-de-seguir", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	//SEMPRE RODAR
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

// SeguirUsuario para de seguir um usuario
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 18, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	// fazendo a chamada para API
	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	//SEMPRE RODAR
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

// EditarUsuario chama a API para editar o usuario
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	// vamos fazer o mesmo papel do POSTMAN, criar uma requisicao e enviar
	// aqui estamos pegando os campos do formulario
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	// detectar id do usuarioLogado
	cookie, _ := cookies.Ler(r)
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	// fazendo a chamada para API
	fmt.Println("chamar api")
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioLogadoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	//SEMPRE RODAR
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

// AtualizarSenhaDoUsuario chama a API para atualizar senha do usuario
func AtualizarSenhaDoUsuario(w http.ResponseWriter, r *http.Request) {
	// vamos fazer o mesmo papel do POSTMAN, criar uma requisicao e enviar
	// aqui estamos pegando os campos do formulario
	r.ParseForm()

	senhas, erro := json.Marshal(map[string]string{
		"atual": r.FormValue("atual"),
		"nova":  r.FormValue("nova"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	// detectar id do usuarioLogado
	cookie, _ := cookies.Ler(r)
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	// fazendo a chamada para API
	fmt.Println("chamar api")
	url := fmt.Sprintf("%s/usuarios/%d/atualizar-senha", config.APIURL, usuarioLogadoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(senhas))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	//SEMPRE RODAR
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

// DeletarUsuario chama a API para excluir o usuario
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	// detectar id do usuarioLogado
	cookie, _ := cookies.Ler(r)
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	// fazendo a chamada para API
	fmt.Println("chamar api")
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioLogadoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodDelete, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	//SEMPRE RODAR
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}
