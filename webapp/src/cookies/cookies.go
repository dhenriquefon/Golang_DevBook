package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Configurar utiliza as variaveis de ambiente para criacao do securecookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// responsewriter eh quem escreve os dados no browser

//Salvar registrar as informacoes de autenticacao
func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}

	//codificar os dados para salvar no cookie (primeiro parametro do Encode eh o nome do cookie, pra quem estou passando meus dados)
	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

//Ler retorna os valores armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) {
	//ler o cookie do browser
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	// decodificar
	valores := make(map[string]string)
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}
