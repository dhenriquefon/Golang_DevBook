package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/respostas"
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
