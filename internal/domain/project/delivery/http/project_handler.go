package http

import (
	"net/http"
	"strconv"

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

type responseProjects struct {
	Message  string                `json:"message"`
	Projects []domain.ProjectModel `json:"projects"`
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
	r.GET("/project/:id", mid.AuthRouteMiddleware(), handler.GetProject)
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

func (a *projectHandler) GetProject(r *gin.Context) {
	id := r.Param("id")

	idProject, errCon := strconv.Atoi(id)
	if errCon != nil {
		r.JSON(http.StatusBadRequest, responseError{Message: errCon.Error()})
		return
	}

	projects, errProjects := a.projectUsecase.Get(r, idProject)
	if errProjects != nil {
		r.JSON(http.StatusBadRequest, responseError{Message: errProjects.Error()})
		return
	}

	res := responseProjects{
		Projects: projects,
		Message:  "get project success",
	}

	r.JSON(http.StatusOK, res)
}
