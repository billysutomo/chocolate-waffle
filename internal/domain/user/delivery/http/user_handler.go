package http

import (
	"net/http"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/middleware"

	"github.com/gin-gonic/gin"
)

// responseError responseError
type responseError struct {
	Message string `json:"message"`
}

// requestRegister requestRegister
type requestRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// requestLogin request
type requestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// requestRefreshToken requestRefreshToken
type requestRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

// userHandler article handler
type userHandler struct {
	UserUsecase domain.UserUsecase
}

// NewUserHandler url handler
func NewUserHandler(r *gin.Engine, mid *middleware.MainMiddleware, do domain.UserUsecase) {
	handler := &userHandler{
		UserUsecase: do,
	}

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("/refresh-token", handler.RefreshToken)
}

// Register Register
func (a *userHandler) Register(r *gin.Context) {
	var reqRegister requestRegister
	r.BindJSON(&reqRegister)

	_, err := a.UserUsecase.CreateUser(r, reqRegister.Name, reqRegister.Email, reqRegister.Password)

	if err != nil {
		r.JSON(http.StatusUnauthorized, responseError{Message: err.Error()})
		return
	}
	r.JSON(http.StatusCreated, map[string]string{
		"message": "user created",
	})
}

// Login login
func (a *userHandler) Login(r *gin.Context) {
	var reqLogin requestLogin
	r.BindJSON(&reqLogin)

	token, refreshToken, err := a.UserUsecase.Login(r, reqLogin.Email, reqLogin.Password)

	if err != nil {
		r.JSON(http.StatusUnauthorized, responseError{Message: err.Error()})
		return
	}
	r.JSON(http.StatusOK, map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
	})
}

// RefreshToken RefreshToken
func (a *userHandler) RefreshToken(r *gin.Context) {
	var reqtRefreshToken requestRefreshToken
	r.BindJSON(&reqtRefreshToken)

	token, refreshToken, err := a.UserUsecase.RefreshToken(r, reqtRefreshToken.RefreshToken)

	if err != nil {
		r.JSON(http.StatusUnauthorized, responseError{Message: err.Error()})
		return
	}
	r.JSON(http.StatusOK, map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
	})
}
