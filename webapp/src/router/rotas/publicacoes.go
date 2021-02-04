package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublicacoes = []Rota{
	{
		URI:               "/publicacoes",
		Metodo:            http.MethodPost,
		Funcao:            controllers.CriarPublicacao,
		RequerAutorizacao: true,
	},
	{
		URI:               "/publicacoes/{publicacaoId}/curtir",
		Metodo:            http.MethodPost,
		Funcao:            controllers.CurtirPublicacao,
		RequerAutorizacao: true,
	},
	{
		URI:               "/publicacoes/{publicacaoId}/descurtir",
		Metodo:            http.MethodPost,
		Funcao:            controllers.DescurtirPublicacao,
		RequerAutorizacao: true,
	},
	{
		URI:               "/publicacoes/{publicacaoId}/atualizar",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarPaginaDeAtualizacaoDePublicacao,
		RequerAutorizacao: true,
	},
	{
		URI:               "/publicacoes/{publicacaoId}",
		Metodo:            http.MethodPut,
		Funcao:            controllers.EditarPublicacao,
		RequerAutorizacao: true,
	},
	{
		URI:               "/publicacoes/{publicacaoId}",
		Metodo:            http.MethodDelete,
		Funcao:            controllers.DeletarPublicacao,
		RequerAutorizacao: true,
	},
}
