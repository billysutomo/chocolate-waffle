package http

import (
	"net/http"

	domain "github.com/billysutomo/chocolate-waffle/internal/domain/url/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ResponseError response error
type ResponseError struct {
	Message string `json:"message"`
}

// URLHandler article handler
type URLHandler struct {
	URLUsecase domain.URLUsecase
}

// NewURLHandler url handler
func NewURLHandler(r *gin.Engine, do domain.URLUsecase) {
	handler := &URLHandler{
		URLUsecase: do,
	}

	r.POST("/url", handler.CreateURL)
}

func isRequestValid(m *domain.URL) (bool, error) {
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
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError

	}
}

// CreateURL create url
func (a *URLHandler) CreateURL(r *gin.Context) {
	var url domain.URL
	r.BindJSON(&url)
	// if ok, err := isRequestValid(&url); !ok {
	// 	r.JSON(http.StatusUnprocessableEntity, err.Error())
	// }

	// ctx := r.Request().Context()
	// err = a.
	r.JSON(http.StatusCreated, url)
}
