package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Cria a request com timeout
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatalf("Erro criando a request: %v", err)
	}

	// Executa a request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Erro ao executar a request: %v", err)
	}
	defer res.Body.Close()

	// Mostra a resposta no stdout
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatalf("Erro lendo o body: %v", err)
	}
}
