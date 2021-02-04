package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL url de conexao com api
	APIURL = ""
	//Porta onde a aplicacao web esta rodando
	Porta = 0
	//HashKey utilizada para autenticacar o cookie
	HashKey []byte
	//BlockKey usada para criptografar os dados do cookie
	BlockKey []byte
)

//Carregar inicializa as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
