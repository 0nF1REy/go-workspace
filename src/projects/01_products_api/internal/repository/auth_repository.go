package repository

import (
	"database/sql"
	"errors"

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

func (ar *AuthRepository) CreateUser(username, passwordHash string) (model.User, error) {
	query := `
		INSERT INTO products.app_user (username, password_hash)
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
		FROM products.app_user
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
