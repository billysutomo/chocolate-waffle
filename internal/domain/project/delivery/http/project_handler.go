package http

import (
	"net/http"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/middleware"
	"github.com/gin-gonic/gin"
)

type responseError struct {
	Message string `json:"message"`
}

type requestProject struct {
	URL            string `json:"url"`
	ProfilePicture string `json:"profile_picture"`
	Title          string `json:"title"`
	Description    string `json:"description"`
}

type projectHandler struct {
	projectUsecase domain.ProjectUsecase
}

// NewProjectHandler NewProjectHandler
func NewProjectHandler(r *gin.Engine, mid *middleware.MainMiddleware, do domain.ProjectUsecase) {
	handler := &projectHandler{
		projectUsecase: do,
	}

	r.POST("/project", mid.AuthRouteMiddleware(), handler.CreateProject)
}

// CreateProject CreateProject
func (a *projectHandler) CreateProject(r *gin.Context) {
	id := r.GetInt("id")

	var reqProject requestProject
	r.BindJSON(&reqProject)

	_, err := a.projectUsecase.CreateProject(
		r,
		id,
		reqProject.URL,
		reqProject.ProfilePicture,
		reqProject.Title,
		reqProject.Description,
	)

	if err != nil {
		r.JSON(http.StatusBadRequest, responseError{Message: err.Error()})
		return
	}
	r.JSON(http.StatusOK, map[string]string{
		"message": "project created",
	})
}
