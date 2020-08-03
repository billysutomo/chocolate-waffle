package http

import (
	"log"
	"net/http"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain/url/model"
	"github.com/dgrijalva/jwt-go"

	_authMiddleware "github.com/billysutomo/chocolate-waffle/internal/domain/auth/controller/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ResponseError response error
type ResponseError struct {
	Message string `json:"message"`
}

// RequestLogin request
type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// URLHandler article handler
type URLHandler struct {
}

// NewURLHandler url handler
func NewURLHandler(r *gin.Engine) {
	handler := &URLHandler{}
	middle := _authMiddleware.InitMiddleware()

	r.POST("/login", handler.Login)
	r.POST("/private", middle.AuthRouteMiddleware(), handler.Private)
}

func isRequestValid(m *model.URL) (bool, error) {
	var v *validator.Validate
	v = validator.New()
	err := v.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case model.ErrInternalServerError:
		return http.StatusInternalServerError
	case model.ErrNotFound:
		return http.StatusNotFound
	case model.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError

	}
}

// Login login
func (a *URLHandler) Login(r *gin.Context) {
	var requestLogin RequestLogin
	r.BindJSON(&requestLogin)
	// Check in your db if the user exists or not
	if requestLogin.Username == "jon" && requestLogin.Password == "password" {
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
			log.Fatal(err)
		}
		r.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	} else {
		r.JSON(http.StatusUnauthorized, ResponseError{Message: "Unauthorized"})
	}
}

// Private Private
func (a *URLHandler) Private(r *gin.Context) {
	r.JSON(http.StatusOK, ResponseError{Message: "Hola"})
}
