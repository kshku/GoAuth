package services

import (
	"context"

	"go_auth/internal/domain"
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
	user := domain.CreateUser(data.Name, data.Email, data.Password)
	if user == nil {
		return nil, ErrValidation
	}

	return &RegisterUserResponse{
		Name: user.Name,
		Email: user.Email,
		Token: "",
	}, nil
}

func LoginUser(ctx context.Context, data *LoginUserData) (*LoginUserResponse, error) {
	return &LoginUserResponse{
		Name: "",
		Email: "",
		Token: "",
	}, nil
}
