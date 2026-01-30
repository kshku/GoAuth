package services

import (
	"context"
	"strings"

	"go_auth/internal/db"
)

type RegisterUserData struct {
	Name string
	Email string
	Password string
}

type RegisterUserResponse struct {
	Name string
	Email string
	Token string
}

type LoginUserData struct {
	Email string
	Password string
}

type LoginUserResponse RegisterUserResponse

func RegisterNewUser(ctx context.Context, data *RegisterUserData) (*RegisterUserResponse, error) {
	if data.Name == "" || data.Email == "" || data.Password == "" {
		return nil, ErrValidation
	}

	data.Name = strings.TrimSpace(data.Name)
	data.Email = strings.ToLower(strings.TrimSpace(data.Email))

	if exists, err := db.UserExists(ctx, data.Email); err != nil {
		return nil, ErrInternal
	} else if exists {
		return nil, ErrUserExists
	}

	passwordHash, err := HashPassword(data.Password)
	if err != nil {
		return nil, ErrPasswordTooLong
	}

	user, err := db.CreateUser(ctx, data.Name, data.Email, passwordHash)
	if err != nil {
		return nil, ErrInternal
	}

	token := ""

	return &RegisterUserResponse{
		Name: user.Name,
		Email: user.Email,
		Token: token,
	}, nil
}

func LoginUser(ctx context.Context, data *LoginUserData) (*LoginUserResponse, error) {
	if data.Email == "" || data.Password == "" {
		return nil, ErrValidation
	}

	data.Email = strings.ToLower(strings.TrimSpace(data.Email))

	user, err := db.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	token:= ""

	return &LoginUserResponse{
		Name: user.Name,
		Email: user.Email,
		Token: token,
	}, nil
}
