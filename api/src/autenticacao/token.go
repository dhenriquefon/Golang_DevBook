package autenticacao

import (
	"devbook-api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken cria um token para o usuario, com as permisoes do usuario
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}

	permissoes["authorized"] = true

	// definindo data que o token expira
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()

	// dono do token
	permissoes["usuariodId"] = usuarioID

	// secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	// assinar o token (tamo chumbando o secret aqui)
	return token.SignedString([]byte(config.SecretKey))

}

// ValidarToken se o token passado na requisicao eh valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	fmt.Println(tokenString)
	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil {
		fmt.Println("token invalido")
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("token invalido")
		return nil
	}

	return errors.New("Token inválido")
}

// ExtrairUsuarioID retorna o ID do usuario que esta salvo no TOKEN
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil {
		return 0, erro
	}

	fmt.Println("ExtrairUsuarioId")
	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuariodId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New("Token inválido")

}

func extrairToken(r *http.Request) string {
	//extraindo o token
	token := r.Header.Get("Authorization")

	// precisamos fazer o split, pois o valor que pode chegar aqui eh Bearer <token>
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
