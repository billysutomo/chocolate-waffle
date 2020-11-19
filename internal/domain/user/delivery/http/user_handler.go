package http

import (
	"net/http"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/middleware"

	"github.com/gin-gonic/gin"
)

// ResponseError ResponseError
type ResponseError struct {
	Message string `json:"message"`
}

// RequestRegister RequestRegister
type RequestRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RequestLogin request
type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RequestRefreshToken RequestRefreshToken
type RequestRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

// UserHandler article handler
type UserHandler struct {
	UserUsecase domain.UserUsecase
}

// NewUserHandler url handler
func NewUserHandler(r *gin.Engine, mid *middleware.MainMiddleware, do domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: do,
	}

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("/refresh-token", handler.RefreshToken)
}

// Register Register
func (a *UserHandler) Register(r *gin.Context) {
	var requestRegister RequestRegister
	r.BindJSON(&requestRegister)

	_, err := a.UserUsecase.CreateUser(r, requestRegister.Name, requestRegister.Email, requestRegister.Password)

	if err != nil {
		r.JSON(http.StatusUnauthorized, ResponseError{Message: err.Error()})
		return
	}
	r.JSON(http.StatusCreated, map[string]string{
		"message": "user created",
	})
}

// Login login
func (a *UserHandler) Login(r *gin.Context) {
	var requestLogin RequestLogin
	r.BindJSON(&requestLogin)

	token, refreshToken, err := a.UserUsecase.Login(r, requestLogin.Email, requestLogin.Password)

	if err != nil {
		r.JSON(http.StatusUnauthorized, ResponseError{Message: err.Error()})
		return
	}
	r.JSON(http.StatusOK, map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
	})
}

// RefreshToken RefreshToken
func (a *UserHandler) RefreshToken(r *gin.Context) {
	var requestRefreshToken RequestRefreshToken
	r.BindJSON(&requestRefreshToken)

	token, refreshToken, err := a.UserUsecase.RefreshToken(r, requestRefreshToken.RefreshToken)

	if err != nil {
		r.JSON(http.StatusUnauthorized, ResponseError{Message: err.Error()})
		return
	}
	r.JSON(http.StatusOK, map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
	})
}
