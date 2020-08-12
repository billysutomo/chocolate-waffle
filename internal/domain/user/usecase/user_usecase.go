package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

// NewUserUsecase NewUserUsecase
func NewUserUsecase(a domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: a,
	}
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

func hashAndSalt(password string) string {

	pwd := []byte(password)

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, password string) bool {
	plainPwd := []byte(password)
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (a *userUsecase) Login(ctx context.Context, email string, password string) (string, string, error) {
	user, err := a.userRepo.GetUserByEmail(ctx, email)

	if err != nil || user.ID == 0 {
		return "", "", errors.New("Email wrong")
	}

	pwdCheck := comparePasswords(user.Password, password)

	if !pwdCheck {
		return "", "", errors.New("Password wrong")
	}

	token, refreshToken, err := generateTokenPair(user.Name, user.ID)
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
}

func (a *userUsecase) RefreshToken(c context.Context, refreshToken string) (string, string, error) {

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

// RegisterUser RegisterUser
func (a *userUsecase) CreateUser(ctx context.Context, name string, email string, password string) (bool, error) {

	hashPassword := hashAndSalt(password)

	_, err := a.userRepo.CreateUser(ctx, name, email, hashPassword)
	if err != nil {
		return false, err
	}
	return true, nil
}
