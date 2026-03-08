package routes

import (
    "fmt"
    "net/http"
)

func SetupRoutes() {
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "pong")
    })
}
