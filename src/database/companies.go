package database

import (
	"fmt"
	"yawoenapi/src/models"

	"github.com/google/uuid"
)

func InsertCompany(companies []models.CompanyInsert) (string, error) {
	//Conecta com o banco de dados
	db, err := ConnectDatabase()
	if err != nil {
		errDescription := "Erro na conexão do banco!"
		return errDescription, err
	}
	defer db.Close()

	//Cria um statement para inserção
	statement, err := db.Prepare("INSERT INTO companies (id, name, zip) VALUES ($1, $2, $3)")
	if err != nil {
		errDescription := "Erro na criação do statement de insert!"
		return errDescription, err
	}
	defer statement.Close()

	//Cria um statement para verificação de duplicatas
	statementAux, err := db.Prepare("SELECT id FROM companies WHERE name = $1 AND zip = $2")
	if err != nil {
		errDescription := "Erro na criação do statement de verificação!"
		return errDescription, err
	}
	defer statementAux.Close()

	//Executa o statement no slice, dando um UUID à companhia no processo
	for _, company := range companies {

		//Verifica se há uma duplicata no banco
		result, err := statementAux.Query(company.Name, company.Zip)
		if err != nil {
			errDescription := "Erro na verificação de duplicata!"
			return errDescription, err
		}
		if result.Next() {
			continue
		}

		//Caso não exista duplicata, tenta a inserção
		_, err = statement.Exec(uuid.New(), company.Name, company.Zip)
		if err != nil {
			errDescription := "Erro na execução do statement de update!"
			return errDescription, err
		}
	}
	return "", nil
}

func AddWebsite() (string, error) {
	//Conecta com o banco de dados
	db, err := ConnectDatabase()
	if err != nil {
		errDescription := "Erro na conexão do banco!"
		return errDescription, err
	}
	defer db.Close()

	//Checa se existe uma coluna chamada website
	result, err := db.Query("SELECT column_name FROM information_schema.columns WHERE table_name='companies' and column_name='website';")
	if err != nil {
		errDescription := "Erro na verificação da coluna!"
		return errDescription, err
	}
	defer result.Close()

	//Se não houver coluna website, ela é criada
	if !result.Next() {
		db.Exec("ALTER TABLE companies ADD website VARCHAR(255)")
		fmt.Println("Coluna website criada!")

		//Adiciona um valor vazio para todas as companhias
		_, err = db.Exec("UPDATE companies SET website = ' '")
		if err != nil {
			errDescription := "Erro na criação do statement de update!"
			return errDescription, err
		}
	}

	return "", nil
}

func UpdateCompany(companies []models.Company) (string, error) {
	//Conecta com o banco de dados
	db, err := ConnectDatabase()
	if err != nil {
		errDescription := "Erro na conexão do banco!"
		return errDescription, err
	}
	defer db.Close()

	//Cria um statement para update
	statementUpdate, err := db.Prepare("UPDATE companies SET website = $1 WHERE name = $2 AND zip = $3")
	if err != nil {
		errDescription := "Erro na criação do statement de update!"
		return errDescription, err
	}
	defer statementUpdate.Close()

	//Cria um statement para verificação se já existe
	statementVerification, err := db.Prepare("SELECT id FROM companies WHERE name = $1 AND zip = $2")
	if err != nil {
		errDescription := "Erro na criação do statement de verificação!"
		return errDescription, err
	}
	defer statementVerification.Close()

	//Cria um statement para inserção
	statementInsert, err := db.Prepare("INSERT INTO companies (id, name, zip, website) VALUES ($1, $2, $3, $4)")
	if err != nil {
		errDescription := "Erro na criação do statement de inserção!"
		return errDescription, err
	}
	defer statementInsert.Close()

	//Executa o statement no slice
	for _, company := range companies {

		//Verifica se já existe no banco
		result, err := statementVerification.Query(company.Name, company.Zip)
		if err != nil {
			errDescription := "Erro na execução do statement de existencia no banco!"
			return errDescription, err
		}

		if result.Next() {
			//Se já existe no banco, dá update
			_, err = statementUpdate.Exec(company.WebSite, company.Name, company.Zip)
			if err != nil {
				errDescription := "Erro na execução do statement de update!"
				return errDescription, err
			}
		} else {
			//Caso não exista, tenta a inserção
			_, err = statementInsert.Exec(uuid.New(), company.Name, company.Zip, company.WebSite)
			if err != nil {
				errDescription := "Erro na execução do statement de inserção!"
				return errDescription, err
			}
		}
	}
	return "", nil
}

func GetCompany(company models.Company) ([]models.Company, string, error) {

	var companies []models.Company
	//Conecta com o banco de dados
	db, err := ConnectDatabase()
	if err != nil {
		errDescription := "Erro na conexão do banco!"
		return companies, errDescription, err
	}
	defer db.Close()

	//Cria um statement para select
	statement, err := db.Prepare("SELECT * FROM companies WHERE name LIKE '%' || $1 || '%' AND zip = $2")
	if err != nil {
		errDescription := "Erro na criação do statement!"
		return companies, errDescription, err
	}
	defer statement.Close()

	//Executa o statement de query
	result, err := statement.Query(company.Name, company.Zip)
	if err != nil {
		errDescription := "Erro na execução do statement!"
		return companies, errDescription, err
	}

	//Pega as linha do SELECT, pois pode haver coincidencia no nome e no zip code, afinal se utiliza apenas parte do nome
	for result.Next() {
		var newCompany models.Company

		if err = result.Scan(&newCompany.Id, &newCompany.Name, &newCompany.Zip, &newCompany.WebSite); err != nil {
			errDescription := "Erro ao escanear usuário!"
			return companies, errDescription, err
		}
		companies = append(companies, newCompany)
	}

	return companies, "", nil
}

//Inserção única para testes
func InsertCompanySingle(company models.CompanyInsert) string {

	db, err := ConnectDatabase()
	if err != nil {
		fmt.Println("Erro na conexão do banco! : " + err.Error())
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO companies (id, name, zip) VALUES ($1, $2, $3)")
	if err != nil {
		fmt.Println("Erro na criação do statement de inserção! : " + err.Error())
	}
	defer statement.Close()

	companyID := uuid.New()

	_, err = statement.Exec(companyID, company.Name, company.Zip)
	if err != nil {
		fmt.Println("Erro na execução do statement! : " + err.Error())
	}

	return companyID.String()
}

//Deleção única para testes
func DeleteCompanySingle(companyID string) {

	db, err := ConnectDatabase()
	if err != nil {
		fmt.Println("Erro na conexão do banco! : " + err.Error())
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM companies WHERE id = $1")
	if err != nil {
		fmt.Println("Erro na criação do statement de deleção! : " + err.Error())
	}
	defer statement.Close()

	_, err = statement.Exec(companyID)
	if err != nil {
		fmt.Println("Erro na execução do statement! : " + err.Error())
	}
}

//Para fins de testes
func RemoveWebsite() {
	db, err := ConnectDatabase()
	if err != nil {
		fmt.Println("Erro na conexão do banco! : " + err.Error())
	}
	defer db.Close()

	_, err = db.Exec("ALTER TABLE companies DROP COLUMN website")
	if err != nil {
		fmt.Println("Erro na criação do statement de update! : " + err.Error())
	}
}
