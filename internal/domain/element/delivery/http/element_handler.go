package http

import (
	"net/http"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/middleware"
	"github.com/gin-gonic/gin"
)

// requestCreateElement requestCreateElement
type requestCreateElement struct {
	IDProject int    `json:"id_project"`
	Ordernum  int    `json:"ordernum"`
	Type      string `json:"type"`
	Body      string `json:"body"`
}

// ElementHandler ElementHandler
type elementHandler struct {
	ElementUsecase domain.ElementUsecase
}

// NewElementHandler NewElementHandler
func NewElementHandler(r *gin.Engine, mid *middleware.MainMiddleware, do domain.ElementUsecase) {
	handler := &elementHandler{
		ElementUsecase: do,
	}

	r.POST("/element", handler.CreateElement)
}

// CreateElement CreateElement
func (a *elementHandler) CreateElement(r *gin.Context) {
	var element requestCreateElement
	errBinding := r.BindJSON(&element)

	if errBinding != nil {
		r.JSON(http.StatusBadRequest, map[string]string{
			"message": errBinding.Error(),
		})
	}

	_, errCreate := a.ElementUsecase.CreateElement(r, element.IDProject, element.Ordernum, element.Type, element.Body)

	if errCreate != nil {
		r.JSON(http.StatusBadRequest, map[string]string{
			"message": errCreate.Error(),
		})
	}

	r.JSON(http.StatusCreated, map[string]string{
		"message": "element created",
	})
}
