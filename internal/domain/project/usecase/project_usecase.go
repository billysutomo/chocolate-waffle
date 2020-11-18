package usecase

import (
	"context"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/domain/project/repository"
	"go.uber.org/zap"
)

type projectUsecase struct {
	projectRepo repository.ProjectRepository
	logger      *zap.Logger
}

// NewProjectUsecase NewProjectUsecase
func NewProjectUsecase(a repository.ProjectRepository, logger *zap.Logger) domain.ProjectUsecase {
	return &projectUsecase{
		projectRepo: a,
		logger:      logger,
	}
}

// CreateProject CreateProject
func (a *projectUsecase) CreateProject(
	ctx context.Context,
	idUser int,
	url string,
	profilePicture string,
	title string,
	description string,
) (bool, error) {
	project := repository.ProjectModel{
		IDUser:         idUser,
		URL:            url,
		ProfilePicture: profilePicture,
		Title:          title,
		Description:    description,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	err := a.projectRepo.CreateProject(ctx, project)
	if err != nil {
		return false, err
	}
	return true, nil
}
