package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/models"
)

func Update(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo

	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := models.Update(id, todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rows == 0 {
		http.Error(w, "todo não encontrado", http.StatusNotFound)
		return
	}

	todo.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
