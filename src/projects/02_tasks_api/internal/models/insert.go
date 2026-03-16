package models

import (
	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/db"
)

func Insert(todo Todo) (id int64, err error) {

	sql := `
	INSERT INTO todos (title, description, done)
	VALUES ($1, $2, $3)
	RETURNING id
`

	err = db.DB.QueryRow(
		sql,
		todo.Title,
		todo.Description,
		todo.Done,
	).Scan(&id)

	return
}
