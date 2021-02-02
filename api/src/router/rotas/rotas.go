package rotas

import (
	"devbook-api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota : objeto que armazena a estrutura da rota, obrigatoria para qualquer rota na aplicacao
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

//Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	// configurando todas as rotas de usuario
	rotas := rotasUsuarios
	rotas = append(rotas, rotasLogin)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			//r.HandleFunc(rota.URI, middlewares.Autenticar(rota.Funcao)).Methods(rota.Metodo)
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	return r
}
