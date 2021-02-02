package main

import (
	"devbook-api/src/config"
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func inicializaParametros() {
	config.Carregar()
	fmt.Println(config.StringConexaoBanco)
	fmt.Println(config.Porta)
}

// gerando secret da aplicacao
//func init() {
//	chave := make([]byte, 64)
//	fmt.Println(chave)
//
//	if _, erro := rand.Read(chave); erro != nil {
//		log.Fatal(erro)
//	}
//
//	fmt.Println(chave)
//
//	stringBase64 := base64.StdEncoding.EncodeToString(chave)
//	fmt.Println(stringBase64)
//}

func main() {
	fmt.Println("Rodando API!")
	fmt.Println("-------------------")
	fmt.Println("Inicializando parametros")
	inicializaParametros()

	r := router.Gerar()
	fmt.Println(config.SecretKey)

	fmt.Println("-------------------")
	fmt.Printf("Escutando na porta :%d", config.Porta)
	// criando meu servidor http ouvindo na porta 5000
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
