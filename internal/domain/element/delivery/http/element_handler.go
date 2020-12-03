package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/middleware"
	"github.com/gin-gonic/gin"
)

type responseElement struct {
	Elements []domain.ElementModel `json:"elements"`
	Message  string                `json:"message"`
}

type requestCreateElement struct {
	IDProject int    `json:"id_project"`
	Ordernum  int    `json:"ordernum"`
	Type      string `json:"type"`
	Body      string `json:"body"`
}

type elementHandler struct {
	ElementUsecase domain.ElementUsecase
}

// NewElementHandler NewElementHandler
func NewElementHandler(r *gin.Engine, mid *middleware.MainMiddleware, do domain.ElementUsecase) {
	handler := &elementHandler{
		ElementUsecase: do,
	}

	r.POST("/element", handler.CreateElement)
	r.GET("/element/:idProject", handler.GetElements)
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

	_, errCreate := a.ElementUsecase.CreateElement(
		r,
		element.IDProject,
		element.Ordernum,
		element.Type,
		element.Body,
		time.Now(),
		time.Now(),
	)

	if errCreate != nil {
		r.JSON(http.StatusBadRequest, map[string]string{
			"message": errCreate.Error(),
		})
	}

	r.JSON(http.StatusCreated, map[string]string{
		"message": "element created",
	})
}

func (a *elementHandler) GetElements(r *gin.Context) {
	idProject := r.Param("idProject")

	i, errCon := strconv.Atoi(idProject)

	if errCon != nil {
		r.JSON(http.StatusBadRequest, map[string]string{
			"message": errCon.Error(),
		})
	}

	elements, err := a.ElementUsecase.GetElementsByIDProject(r, i)

	if err != nil {
		r.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	res := responseElement{
		Elements: elements,
		Message:  "Get elements success",
	}

	r.JSON(http.StatusCreated, res)

}
