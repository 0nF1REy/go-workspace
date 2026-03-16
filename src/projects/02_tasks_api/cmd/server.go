package tasks_api

import (
	"log"

	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/configs"
	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/db"
	"github.com/joho/godotenv"
)

func RunAPI() {
	godotenv.Load()

	if err := configs.Load(); err != nil {
		log.Fatalf("Falha ao carregar configuração: %v", err)
	}
	if err := db.Init(); err != nil {
		log.Fatalf("Falha ao inicializar banco de dados: %v", err)
	}
}
