package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Servidor iniciado na porta 8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Iniciou a request.")
	defer log.Println("Finalizou a request.")

	select {
	case <-time.After(5 * time.Second):
		msg := "Requisição processada com sucesso."
		log.Println(msg)
		w.Write([]byte(msg))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente.")
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)
	}
}
