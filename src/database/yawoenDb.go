package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Dados para a string de conexão
const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "mypassword"
	dbname   = "yawoendb"
)

//Função para garantir que a tabela Companies esteja criada
func InitDatabase() {
	db, err := ConnectDatabase()
	if err != nil {
		fmt.Println("Erro na conexão do banco! : " + err.Error())
	}
	defer db.Close()

	//Tenta criar a tabela caso não exista no banco de dados
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS companies(id VARCHAR(255) NOT NULL, name VARCHAR(255) NOT NULL, zip VARCHAR(255) NOT NULL, PRIMARY KEY(id))")
	if err != nil {
		fmt.Println("Erro na criação de tabela! : " + err.Error())
	}

}

//Função para se conectar ao banco de dados
func ConnectDatabase() (*sql.DB, error) {
	//String de conexão
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Armazena o banco aberto em db
	db, err := sql.Open("postgres", psqlInfo)

	//Caso haja erro na abertura do banco
	if err != nil {
		return nil, err
	}

	//Apesar do banco estar aberto, ping pode não obter resposta
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	//Banco aberto :)
	return db, nil
}
