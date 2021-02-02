package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexaoBanco eh a string de conexao com o banco de dados
	StringConexaoBanco = ""

	//Porta indica a porta que a aplicacao esta rodando
	Porta = 0

	//Chave para assinar o token
	SecretKey []byte
)

// Carrega vai inicializar as variaveis de ambiente
func Carregar() {
	// vamos usar o pacote go get github.com/joho/godotenv para ler os parametros configurados no arqivo .env na raiz do projeto
	var erro error

	// godotenv carrega as variaveis de ambient para o objeto 'os'
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	// le porta api
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		// tratamento de erro jogando para porta padrao
		Porta = 9000
		fmt.Println("Falha ao ler configuracao API_PORT aplicacao iniciada na porta 9000 default")
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"))

	//Secret Key para ser usado nas requisicoes por TOKEY
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
