package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/model"
	"github.com/lib/pq"
)

var ErrDuplicateUser = errors.New("duplicate user")

type AuthRepository struct {
	Connection *sql.DB
}

func NewAuthRepository(connection *sql.DB) AuthRepository {
	return AuthRepository{Connection: connection}
}

func (ar *AuthRepository) EnsureSchema() error {
	query := `
		CREATE TABLE IF NOT EXISTS app_user (
			id SERIAL PRIMARY KEY,
			username VARCHAR(150) UNIQUE NOT NULL,
			password_hash VARCHAR(100) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		)
	`

	if _, err := ar.Connection.Exec(query); err != nil {
		return fmt.Errorf("failed to ensure auth schema: %w", err)
	}

	return nil
}

func (ar *AuthRepository) CreateUser(username, passwordHash string) (model.User, error) {
	query := `
		INSERT INTO app_user (username, password_hash)
		VALUES ($1, $2)
		RETURNING id, username, password_hash, created_at
	`

	var user model.User
	err := ar.Connection.QueryRow(query, username, passwordHash).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			return model.User{}, ErrDuplicateUser
		}
		return model.User{}, err
	}

	return user, nil
}

func (ar *AuthRepository) GetUserByUsername(username string) (model.User, error) {
	query := `
		SELECT id, username, password_hash, created_at
		FROM app_user
		WHERE username = $1
	`

	var user model.User
	err := ar.Connection.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
