package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"yawoenapi/src/database"
	"yawoenapi/src/models"

	"github.com/gorilla/mux"
)

func InitRoute() {
	router := mux.NewRouter()
	fmt.Println("Listening :8000")

	//Cria os caminhos da API
	router.HandleFunc("/yawoenapi", updateDatabase).Methods(http.MethodPatch)
	router.HandleFunc("/yawoenapi", returnCompany).Methods(http.MethodGet)

	//Faz o router escutar e responder na porta 8000
	log.Fatal(http.ListenAndServe(":8000", router))
}

func updateDatabase(result http.ResponseWriter, request *http.Request) {
	//Verifica se existe a coluna website antes do update
	errDescription, err := database.AddWebsite()
	if err != nil {
		result.Write([]byte(errDescription + " : " + err.Error()))
		return
	}

	//Executa o update
	errDescription, err = ReadCSVUpdate("./media/q2_clientData.csv")
	if err != nil {
		result.Write([]byte(errDescription + " : " + err.Error()))
		return
	}

	result.WriteHeader(http.StatusOK)
}

func returnCompany(result http.ResponseWriter, request *http.Request) {
	//Lê o corpo da requisição
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result.Write([]byte("Erro ao ler o corpo da requisição : " + err.Error()))
		return
	}

	//Armazena em company os dados do JSON
	var company models.Company
	if err = json.Unmarshal(requestBody, &company); err != nil {
		result.Write([]byte("Erro ao converter usuários para JSON : " + err.Error()))
		return
	}

	company.Name = strings.ToUpper(company.Name)

	//Verifica se existe a coluna website antes de adquirir o slice
	errDescription, err := database.AddWebsite()
	if err != nil {
		result.Write([]byte(errDescription + " : " + err.Error()))
		return
	}

	//Adquire o slice de companhias no banco de dados
	companies, errDescription, err := database.GetCompany(company)
	if err != nil {
		result.Write([]byte(errDescription + " : " + err.Error()))
		return
	}

	result.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(result).Encode(companies); err != nil {
		result.Write([]byte("Erro ao converter usuários para JSON : " + err.Error()))
		return
	}
}
