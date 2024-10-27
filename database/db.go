package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/titi0001/Go-Gin-RestAPI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDb() {

	// Carrega as variáveis de ambiente de um arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Recupera as variáveis de ambiente
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSL_MODE")

	// Verifica se todas as variáveis de ambiente necessárias estão definidas
	if host == "" || user == "" || password == "" || dbname == "" || port == "" || sslmode == "" {
		log.Fatal("Alguma(s) variável(s) de ambiente não está(ão) definida(s)")
	}

	// Monta a string de conexão
	strConexao := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode

	DB, err = gorm.Open(postgres.Open(strConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})

}
