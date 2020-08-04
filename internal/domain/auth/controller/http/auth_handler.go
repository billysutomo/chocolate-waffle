package http

import (
	"net/http"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/domain/url/model"
	"github.com/billysutomo/chocolate-waffle/internal/middleware"

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

// RequestRefreshToken RequestRefreshToken
type RequestRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

// AuthHandler article handler
type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

// NewAuthHandler url handler
func NewAuthHandler(r *gin.Engine, mid *middleware.MainMiddleware, do domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: do,
	}

	r.POST("/login", handler.Login)
	r.POST("/refresh-token", handler.RefreshToken)
	r.POST("/private", mid.AuthRouteMiddleware(), handler.Private)
}

// Private Private
func (a *AuthHandler) Private(r *gin.Context) {

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
func (a *AuthHandler) Login(r *gin.Context) {
	var requestLogin RequestLogin
	r.BindJSON(&requestLogin)

	token, refreshToken, err := a.AuthUsecase.Login(r, requestLogin.Username, requestLogin.Password)

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
func (a *AuthHandler) RefreshToken(r *gin.Context) {
	var requestRefreshToken RequestRefreshToken
	r.BindJSON(&requestRefreshToken)

	token, refreshToken, err := a.AuthUsecase.RefreshToken(r, requestRefreshToken.RefreshToken)

	if err != nil {
		r.JSON(http.StatusUnauthorized, ResponseError{Message: err.Error()})
		return
	}
	r.JSON(http.StatusOK, map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
	})
}
