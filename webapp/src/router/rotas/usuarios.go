package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuario = []Rota{
	{
		URI:               "/criar-usuario",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutorizacao: false,
	},
	{
		URI:               "/usuarios",
		Metodo:            http.MethodPost,
		Funcao:            controllers.CriarUsuario,
		RequerAutorizacao: false,
	},
}
