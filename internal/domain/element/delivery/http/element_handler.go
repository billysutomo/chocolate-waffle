package http

import (
	"net/http"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RequestCreateElement RequestCreateElement
type RequestCreateElement struct {
	IDProject int    `json:"id_project"`
	Ordernum  int    `json:"ordernum"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	Icon      string `json:"icon"`
	Title     string `json:"title"`
}

// ElementHandler ElementHandler
type ElementHandler struct {
	ElementUsecase domain.ElementUsecase
}

// NewElementHandler NewElementHandler
func NewElementHandler(r *gin.Engine, mid *middleware.MainMiddleware, do domain.ElementUsecase) {
	handler := &ElementHandler{
		ElementUsecase: do,
	}

	r.POST("/element", handler.CreateElement)
}

// CreateElement CreateElement
func (a *ElementHandler) CreateElement(r *gin.Context) {
	r.JSON(http.StatusCreated, map[string]string{
		"message": "element created",
	})
}
