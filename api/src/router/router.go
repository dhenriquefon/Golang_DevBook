package router

import (
	"devbook-api/src/router/rotas"

	"github.com/gorilla/mux"
)

//Gerar vai retornar um router com as rotas configuradas. Lembrando que o router eh quem vai ter as nossas rotas, ou seja nossas urls
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
