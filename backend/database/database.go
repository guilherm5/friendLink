package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Init() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Erro ao carregar variavel de ambiente para logar no banco de dados", err)
	}
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	if host == "" || database == "" || user == "" || password == "" {
		log.Println("Variaveis de ambiente do banco de dados nao pode ser null")
	}

	connection := fmt.Sprintf(`%s:%s@tcp(%s)/%s`, user, password, host, database)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Println("Erro ao conectar com o banco de dados", err)
	} else {
		log.Println("Sucesso ao realizar conexão com banco de dados")
	}
	return db

}
