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
	{
		URI:               "/buscar-usuarios",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarPaginaDeUsuarios,
		RequerAutorizacao: true,
	},
	{
		URI:               "/usuarios/{usuarioId}",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarPerfilDoUsuario,
		RequerAutorizacao: true,
	},
	{
		URI:               "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:            http.MethodPost,
		Funcao:            controllers.PararDeSeguirUsuario,
		RequerAutorizacao: true,
	},
	{
		URI:               "/usuarios/{usuarioId}/seguir",
		Metodo:            http.MethodPost,
		Funcao:            controllers.SeguirUsuario,
		RequerAutorizacao: true,
	},
	{
		URI:               "/perfil",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarPerfilDoUsuarioLogado,
		RequerAutorizacao: true,
	},
	{
		URI:               "/editar-usuario",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarPaginaDeEdicaoDeUsuario,
		RequerAutorizacao: true,
	},
	{
		URI:               "/editar-usuario",
		Metodo:            http.MethodPut,
		Funcao:            controllers.EditarUsuario,
		RequerAutorizacao: true,
	},
	{
		URI:               "/atualizar-senha",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarPaginaDeAtualizacaoDeSenha,
		RequerAutorizacao: true,
	},
	{
		URI:               "/atualizar-senha",
		Metodo:            http.MethodPost,
		Funcao:            controllers.AtualizarSenhaDoUsuario,
		RequerAutorizacao: true,
	},
	{
		URI:               "/deletar-usuario",
		Metodo:            http.MethodDelete,
		Funcao:            controllers.DeletarUsuario,
		RequerAutorizacao: true,
	},
}
