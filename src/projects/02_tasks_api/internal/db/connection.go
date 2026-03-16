package db

import (
	"database/sql"
	"fmt"

	"github.com/0nF1REy/go-workspace/projects/02_tasks_api/internal/configs"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() error {
	conf := configs.GetDB()
	sc := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Database,
	)

	db, err := sql.Open("postgres", sc)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	DB = db

	return nil
}
