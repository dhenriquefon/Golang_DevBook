package modelos

//DadosAutenticacao estrutura com o formato devolvido para aplicacao apos um log
type DadosAutenticacao struct {
	ID    string `json: "id"`
	Token string `json: "token"`
}
