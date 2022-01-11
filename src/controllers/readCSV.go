package controllers

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"yawoenapi/src/database"
	"yawoenapi/src/models"
)

func readCSV(path string) ([][]string, error) {
	//Abre o arquivo CSV presente no caminho
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	//Cria o reader, separando por ";"
	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1

	//Faz com que o reader leia o arquivo e retorne as linhas em csvLines
	csvLines, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return csvLines, err
}

//Função para o primeiro insert
func ReadCSVInsert(path string) {

	var companies []models.CompanyInsert

	csvCompanies, err := readCSV(path)
	if err != nil {
		fmt.Println("Houve um erro na leitura do aquivo: ", err.Error())
	}

	for csvIndex, csvLine := range csvCompanies {
		//Pula a primeira linha, pois é o cabeçalho
		if csvIndex == 0 {
			continue
		}

		zip := csvLine[1]
		if len(zip) == 5 {
			name := strings.ToUpper(csvLine[0])

			//Insere no slice companies
			companies = append(companies, models.CompanyInsert{Name: name, Zip: zip})
		}
	}
	errDescription, err := database.InsertCompany(companies)
	if err != nil {
		fmt.Println(errDescription, " : ", err.Error())
	}
	fmt.Println("Companhias inseridas com sucesso!---------------------------")
}

//Função para o update com websites
func ReadCSVUpdate(path string) (string, error) {

	var companies []models.Company

	csvCompanies, err := readCSV(path)
	if err != nil {
		errDescription := "Houve um erro na leitura do arquivo!"
		return errDescription, err
	}

	for csvIndex, csvLine := range csvCompanies {
		//Pula a primeira linha, pois é o cabeçalho
		if csvIndex == 0 {
			continue
		}

		zip := csvLine[1]
		if len(zip) == 5 {
			name := strings.ToUpper(csvLine[0])
			website := strings.ToLower(csvLine[2])

			//Insere no slice companies
			companies = append(companies, models.Company{Name: name, Zip: zip, WebSite: website})
		}
	}

	errDescription, err := database.UpdateCompany(companies)

	if err != nil {
		fmt.Println("Companhias atualizadas com sucesso!-------------------------")
	}

	return errDescription, err
}
