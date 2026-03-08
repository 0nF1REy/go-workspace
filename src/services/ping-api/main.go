package main

import (
    "fmt"
    "net/http"

    "github.com/0nF1REy/go-workspace/services/ping-api/routes"
)

func main() {
    fmt.Println("API rodando em http://localhost:8080")
    routes.SetupRoutes()
    http.ListenAndServe(":8080", nil)
}
