package http

import (
	"net/http"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/domain/url/model"

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
	AuthUsecase domain.AuthUsecase
}

// NewURLHandler url handler
func NewURLHandler(r *gin.Engine, do domain.AuthUsecase) {
	handler := &URLHandler{
		AuthUsecase: do,
	}
	// middle := _authMiddleware.InitMiddleware()

	r.POST("/login", handler.Login)
	// r.POST("/private", middle.AuthRouteMiddleware(), handler.Private)
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

	token, err := a.AuthUsecase.Login(r, requestLogin.Username, requestLogin.Password)

	if err != nil {
		r.JSON(http.StatusUnauthorized, ResponseError{Message: err.Error()})
		return
	}
	r.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
