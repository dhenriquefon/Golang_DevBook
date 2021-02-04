package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

//FazerRequisicaoComAutenticacao utilizada para colocar o token na requisicao
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	//1 - criar a requisicao
	//2 - passar o token no cabecalho da requisicao
	//3 - client http para fazer a requisicao de fato
	// lembrando que estamos fazendo uma requisicao sobre uma requisicao, webapp e api

	//1 - criar a requisicao
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}

	//2 - 2 - passar o token no cabecalho da requisicao

	//ler o cookie (ignorando o erro, pois se ja chegou ate aqui ja passou pelo middleware) (lendo cookie da requesta da webapp)
	cookie, _ := cookies.Ler(r)
	// passar dados do token no header da requisicao da api (//Bearer faz parte do nosso token)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	//3 - client http para fazer a requisicao de fato

	//criar o client
	client := &http.Client{}
	//fazer a requisicao com token no header (request passado da funcao para api)
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil
}
