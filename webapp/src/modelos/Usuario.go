package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// Usuario representa uma pessoa utilizando a rede social
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoe"`
}

//BuscarUsuarioCompleto monta o objeto cmopleto de usuarios, chamando a api 4 vezes
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	// criar 4 canais (cada um deles fara uma chamada para API que poderao ser feitas em paralelo)
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioID, r)

	// depois usamos o SELECT para sincronizar as chamadas

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("Erro ao buscar usuario")
			}
			usuario = usuarioCarregado
		case seguidoresCarregado := <-canalSeguidores:
			if seguidoresCarregado == nil {
				return Usuario{}, errors.New("Erro ao buscar seguidores")
			}
			seguidores = seguidoresCarregado
		case seguindoCarregado := <-canalSeguindo:
			if seguindoCarregado == nil {
				return Usuario{}, errors.New("Erro ao buscar quem o usuario segue")
			}
			seguindo = seguindoCarregado
		case publicacoesCarregado := <-canalPublicacoes:
			if publicacoesCarregado == nil {
				return Usuario{}, errors.New("Erro ao buscar publicacoes")
			}
			publicacoes = publicacoesCarregado
		}
	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes
	return usuario, nil
}

//BuscarDadosDoUsuario preenche os dados do usuario
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) {
	// nao preciso passar o CORPO entao ja comeco enviando a url a requisicao para API
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		//devolve no canal um usuario vario
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario

}

//BuscarSeguidores traz todos os seguidores de um determinado usuario
func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	// nao preciso passar o CORPO entao ja comeco enviando a url a requisicao para API
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		//devolve no canal um usuario vario
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		//enviar um slice vazio que nao vai ser do tipo nil (pra nao considerar que deu algum erro)
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidores
}

//BuscarSeguindo traz todos os usuarios que um usuario esta segundo
func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	// nao preciso passar o CORPO entao ja comeco enviando a url a requisicao para API
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		//devolve no canal um usuario vario
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}

	if seguindo == nil {
		//enviar um slice vazio que nao vai ser do tipo nil (pra nao considerar que deu algum erro)
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguindo
}

//BuscarPublicacoes traz todas as publicacoes de um usuario
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioID uint64, r *http.Request) {
	// nao preciso passar o CORPO entao ja comeco enviando a url a requisicao para API
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		//devolve no canal um usuario vario
		canal <- nil
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}

	if publicacoes == nil {
		//enviar um slice vazio que nao vai ser do tipo nil (pra nao considerar que deu algum erro)
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- publicacoes
}
