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

// RequestProject RequestProject
type RequestProject struct {
	URL            string `json:"url"`
	ProfilePicture string `json:"profile_picture"`
	Title          string `json:"title"`
	Description    string `json:"description"`
}

// ProjectHandler ProjectHandler
type ProjectHandler struct {
	projectUsecase domain.ProjectUsecase
}

// NewProjectHandler NewProjectHandler
func NewProjectHandler(r *gin.Engine, mid *middleware.MainMiddleware, do domain.ProjectUsecase) {
	handler := &ProjectHandler{
		projectUsecase: do,
	}

	r.POST("/project", mid.AuthRouteMiddleware(), handler.CreateProject)
}

// CreateProject CreateProject
func (a *ProjectHandler) CreateProject(r *gin.Context) {
	id := r.GetInt("id")

	var requestProject RequestProject
	r.BindJSON(&requestProject)

	_, err := a.projectUsecase.CreateProject(
		r,
		id,
		requestProject.URL,
		requestProject.ProfilePicture,
		requestProject.Title,
		requestProject.Description,
	)

	if err != nil {
		r.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	r.JSON(http.StatusOK, map[string]string{
		"message": "project created",
	})
}
