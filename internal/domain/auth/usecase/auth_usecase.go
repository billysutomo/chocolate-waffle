package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
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

func (a *authUsecase) Login(c context.Context, username string, password string) (string, string, error) {
	if username == "jon" && password == "password" {
		token, refreshToken, err := generateTokenPair("Joe Doe", 1)
		if err != nil {
			return "", "", err
		}
		return token, refreshToken, nil
	}
	return "", "", errors.New("error")
}

func generateTokenPair(name string, userID int) (string, string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID          works too)
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["id"] = userID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, errr := refreshToken.SignedString([]byte("secret"))
	if errr != nil {
		return "", "", errr
	}

	return t, rt, nil
}

func (a *authUsecase) RefreshToken(c context.Context, refreshToken string) (string, string, error) {

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secret"), nil
	})

	if err != nil {
		log.Print(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		name := "Joe Doe"
		id := 1
		if int(claims["id"].(float64)) == 1 {

			newToken, newRefreshToken, err := generateTokenPair(name, id)
			if err != nil {
				return "", "", err
			}
			return newToken, newRefreshToken, nil
		}
		return "", "", errors.New("Unauthorized")
	}
	return "", "", err
}
