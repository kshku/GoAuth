package services

import (
	"os"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
var SECRET string

func InitJWT() error {
	SECRET = os.Getenv("JWT_SECRET")
	if SECRET == "" {
		return errors.New("JWT_SECRET is not set")
	}
	return nil
}


func GenerateJWT(userID int, email string) (string, error) {
	SECRET := os.Getenv("JWT_SECRET")
	if SECRET == "" {
		return "", errors.New("JWT_SECRET is not set")
	}

	// Token expires in 24 hours
	expairationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expairationTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStirng, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}

	return tokenStirng, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return claims, nil
}

