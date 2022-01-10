package database

import (
	"testing"
	"yawoenapi/src/models"
)

type cenarioGetTeste struct {
	name            string
	zip             string
	companiesLength int
}

type cenarioUpdateTeste struct {
	name    string
	zip     string
	website string
	err     error
}

//Inserção de uma companhia para testes
func insertCompany(name string, zip string) string {
	return InsertCompanySingle(models.CompanyInsert{Name: name, Zip: zip})
}

//Deleção de uma companhia para testes
func deleteCompany(id string) {
	DeleteCompanySingle(id)
}

//Espera-se que o UpdateCompany possa atualizar a tabela sem erros
func TestUpdateCompany(t *testing.T) {
	InitDatabase()
	id := insertCompany("TESTE", "TESTE")
	id2 := insertCompany("ETSET", "TESTE")
	AddWebsite()
	defer deleteCompany(id)
	defer deleteCompany(id2)
	defer RemoveWebsite()

	cenarios := []cenarioUpdateTeste{
		{"TESTE", "TESTE", "siteTESTE.com", nil},
		{"ETSET", "ETSET", "siteETSET.com", nil},
	}

	var cenariosTestes []models.Company
	for _, cenario := range cenarios {
		company := models.Company{Name: cenario.name, Zip: cenario.zip, WebSite: cenario.website}
		cenariosTestes = append(cenariosTestes, company)
	}
	_, err := UpdateCompany(cenariosTestes)
	if err != nil {
		t.Errorf("Update retornou erro, esperava nil! Erro: %s", err.Error())
	}
}

//Espera-se que o GetCompany possa retornar uma companhia que contenha parte de name e zip de Company
func TestGetCompany(t *testing.T) {
	InitDatabase()
	id := insertCompany("TESTE", "TESTE")
	id2 := insertCompany("ETSET", "TESTE")
	AddWebsite()
	defer deleteCompany(id)
	defer deleteCompany(id2)
	defer RemoveWebsite()

	cenarios := []cenarioGetTeste{
		{"T", "T", 0},
		{"TESTE", "T", 0},
		{"T", "TESTE", 2},
		{"TESTE", "TESTE", 1},

		{"E", "T", 0},
		{"ETSET", "", 0},
		{"E", "TESTE", 2},
		{"ETSET", "TESTE", 1},
	}

	for _, cenario := range cenarios {
		companies, _, _ := GetCompany(models.Company{Name: cenario.name, Zip: cenario.zip})
		companiesLength := len(companies)
		if cenario.companiesLength != companiesLength {
			t.Errorf("Quantidade de retornos %d é diferente do esperado %d", companiesLength, cenario.companiesLength)
		}
	}

}
