package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/respostas"
)

//FazerLogin valida o usuario para entrar na aplicacao
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue(("email")),
		"senha": r.FormValue(("senha")),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	//fazendo a requisicao para a API, passando usuario no formato JSON
	url := fmt.Sprintf("%s/login", config.APIURL)
	fmt.Println(url)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	//token do usuario esta no corpo da resposta (certo? eh retornado pela api)
	//token, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(response.StatusCode, string(token))

	fmt.Println(response.StatusCode)
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	// converte os dados de JSON para dados Autenticacao para obtermos ID e token do usuario para requisicoes
	var dadosAutenticacao modelos.DadosAutenticacao
	if erro = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	// salvar os dados dentro do cookie
	fmt.Println("Salvarcookie")
	if erro = cookies.Salvar(w, dadosAutenticacao.ID, dadosAutenticacao.Token); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	fmt.Println("OK")
	respostas.JSON(w, http.StatusOK, nil)
}
