package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	Host string
	Port string
	User string
	Pass string
	DB   string
)

func init() {
	// Detecta o diretório do arquivo conn.go
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	// Caminho do .env relativo a este arquivo
	envPath := filepath.Join(basepath, "../../../.env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("Aviso: não foi possível carregar o .env, usando variáveis de ambiente do sistema")
	}

	Host = os.Getenv("POSTGRES_HOST")
	Port = os.Getenv("POSTGRES_PORT")
	User = os.Getenv("POSTGRES_USER")
	Pass = os.Getenv("POSTGRES_PASSWORD")
	DB = os.Getenv("POSTGRES_DB")

	if Host == "" || Port == "" || User == "" || Pass == "" || DB == "" {
		log.Println("Aviso: algumas variáveis do .env não foram carregadas corretamente")
		log.Printf("POSTGRES_HOST = %s\nPOSTGRES_PORT = %s\nPOSTGRES_USER = %s\nPOSTGRES_PASSWORD = %s\nPOSTGRES_DB = %s\n",
			Host, Port, User, Pass, DB)
	}
}

func GetPostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Pass, DB)
}

func ConnectDB() (*sql.DB, error) {
	dsn := GetPostgresDSN()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir conexão: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Erro ao conectar: %w", err)
	}

	fmt.Println("Conexão com o PostgreSQL realizada com sucesso!")
	return db, nil
}
