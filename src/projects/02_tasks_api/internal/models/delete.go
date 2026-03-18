package models

import (
	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/db"
)

func Delete(id int64) (int64, error) {

	result, err := db.DB.Exec(`
		       DELETE FROM tasks.todos
		       WHERE id = $1
	       `, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
