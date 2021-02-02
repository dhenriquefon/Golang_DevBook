package modelos

//Senha representa a nova sneha e atual para alterar a senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
