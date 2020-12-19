package usecase

import (
	"context"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"go.uber.org/zap"
)

type projectUsecase struct {
	projectRepo domain.ProjectRepository
	logger      *zap.Logger
}

// NewProjectUsecase NewProjectUsecase
func NewProjectUsecase(a domain.ProjectRepository, logger *zap.Logger) domain.ProjectUsecase {
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
	project := domain.ProjectModel{
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

func (a *projectUsecase) Get(ctx context.Context, id int) ([]domain.ProjectModel, error) {
	projects, err := a.projectRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
