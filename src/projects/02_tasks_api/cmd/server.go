package tasks_api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/configs"
	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/db"
	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func RunAPI() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if err := configs.Load(); err != nil {
		log.Fatalf("Falha ao carregar configuração: %v", err)
	}

	if err := db.Init(); err != nil {
		log.Fatalf("Falha ao inicializar banco de dados: %v", err)
	}

	r := chi.NewRouter()

	// ----------------------
	// Task API Endpoints
	// ----------------------
	r.Post("/todos", handlers.Create)
	r.Get("/todos", handlers.List)
	r.Get("/todos/{id}", handlers.Get)
	r.Put("/todos/{id}", handlers.Update)
	r.Delete("/todos/{id}", handlers.Delete)

	addr := fmt.Sprintf(":%s", configs.GetServerPort())

	log.Printf("Server running on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, r))
}
