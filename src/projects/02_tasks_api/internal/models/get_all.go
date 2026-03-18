package models

import (
	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/db"
)

func GetAll() (todos []Todo, err error) {

	rows, err := db.DB.Query(`
		       SELECT id, title, description, done
		       FROM tasks.todos
	       `)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {

		var todo Todo

		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Done,
		)
		if err != nil {
			return
		}

		todos = append(todos, todo)
	}

	return
}
