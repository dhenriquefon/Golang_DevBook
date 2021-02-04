package modelos

//DadosAutenticacao armazenda os dados da autenticacao do usuario, token e ID
type DadosAutenticacao struct {
	ID    string `json: "id"`
	Token string `json: "token"`
}
