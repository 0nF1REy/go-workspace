package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/models"
)

func List(w http.ResponseWriter, r *http.Request) {

	todos, err := models.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
