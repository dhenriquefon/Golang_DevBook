package rotas

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

//Rota representa todas as rotas da app web
type Rota struct {
	URI               string
	Metodo            string
	Funcao            func(w http.ResponseWriter, r *http.Request)
	RequerAutorizacao bool
}

//Configurar coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuario...)
	rotas = append(rotas, rotasPublicacoes...)
	rotas = append(rotas, rotaPaginaPrincipal)
	rotas = append(rotas, rotaLogout)

	for _, rota := range rotas {

		if rota.RequerAutorizacao {
			router.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}
	}

	//apontando que nossos arquivos de estilo e de javascript estao nessa pasta
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
