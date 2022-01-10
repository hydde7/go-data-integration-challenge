package models

//Modelo para a primeira inserção
type CompanyInsert struct {
	Name string `json:"name"`
	Zip  string `json:"zip"`
}

//Modelo para as transações no banco de dados
type Company struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Zip     string `json:"zip"`
	WebSite string `json:"website"`
}
