package models

import (
	"database/sql"

	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/db"
)

func Get(id int64) (todo Todo, err error) {

	row := db.DB.QueryRow(`
		SELECT id, title, description, done
		FROM tasks.todos
		WHERE id = $1
	`, id)

	err = row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Done,
	)

	if err == sql.ErrNoRows {
		return Todo{}, sql.ErrNoRows
	}

	return
}
