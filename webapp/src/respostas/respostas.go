package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

//ErroApi representa a resposta de erro da API
type ErroApi struct {
	//Erro representa a mensagem de erro
	Erro string `json: "erro"`
}

// JSON retorna uma resposta em formato JSON para requisoca
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	//essa nao eh a melhor validacao, pois qualquer chamada ajax sempre espera um conteudo na resposta
	//if dados != nil {
	//	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
	//		log.Fatal(erro)
	//	}
	//}

	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}

}

// TratarStatusCodeDeErro trata as requisições com status code 400 ou superior
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroApi
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)

}
