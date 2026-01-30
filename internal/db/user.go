package db

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID int
	Name string
	Email string
	PasswordHash string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(ctx context.Context, name, email, passwordHash string) (*User, error) {
	query := `
		INSERT INTO users (name, email, password_hash, create_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id, name, email, password_hash, create_at, updated_at
	`

	var user User
	err := DB.QueryRowContext(ctx, query, name, email, passwordHash).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func UserExists(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	var exists bool
	if err := DB.QueryRowContext(ctx, query, email).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user User
	err := DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
