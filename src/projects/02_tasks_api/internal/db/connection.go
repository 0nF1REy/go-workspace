package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/configs"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() error {
	conf := configs.GetDB()

	sc := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Database,
	)

	db, err := sql.Open("postgres", sc)
	if err != nil {
		return fmt.Errorf("erro ao abrir conexão com o banco: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			DB = db
			log.Println("Banco de dados conectado com sucesso!")
			return nil
		}

		log.Println("Aguardando o banco de dados ficar disponível...")
		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("Não foi possível conectar ao banco após várias tentativas: %w", err)
}
