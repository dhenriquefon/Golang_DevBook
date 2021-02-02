package middlewares

import (
	"devbook-api/src/autenticacao"
	"devbook-api/src/respostas"
	"log"
	"net/http"
)

// camada que fica entre a requisicao e resposta
// eh muito utilizado quando voce tem uma funcao que deve ser aplicada para todas as rotas

// hext ttp.HandlerFunc == func(w http.ResponseWriter, r *http.Request)

//Autenticar verifica se usuario que faz a requisicao esta autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		//executa a funcao vinda no parametro
		proximaFuncao(w, r)
	}
}

//Logger escreve informacoes da rota no terminal a cada chamada de rota
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}
