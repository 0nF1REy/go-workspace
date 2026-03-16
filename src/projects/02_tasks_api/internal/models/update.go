package models

import (
	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/db"
)

func Update(id int64, todo Todo) (int64, error) {

	result, err := db.DB.Exec(`
		UPDATE todos
		SET title = $1,
		    description = $2,
		    done = $3
		WHERE id = $4
	`,
		todo.Title,
		todo.Description,
		todo.Done,
		id,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
