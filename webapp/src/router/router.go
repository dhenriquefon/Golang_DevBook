package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar gera as rotas da aplicacao
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
