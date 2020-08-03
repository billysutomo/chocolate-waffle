package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/dgrijalva/jwt-go"
)

type authUsecase struct {
}

// NewAuthUsecase NewAuthUsecase
func NewAuthUsecase() domain.AuthUsecase {
	return &authUsecase{}
}

func (a *authUsecase) Login(c context.Context, username string, password string) (string, error) {
	if username == "jon" && password == "password" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		// This is the information which frontend can use
		// The backend can also decode the token and get admin etc.
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Doe"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
		// Generate encoded token and send it as response.
		// The signing string should be secret (a generated UUID          works too)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return "", errors.New("error")
		}
		return t, nil
	}

	return "", errors.New("error")
}
