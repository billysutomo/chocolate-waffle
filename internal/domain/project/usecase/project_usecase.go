package usecase

import (
	"context"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
)

type projectUsecase struct {
	projectRepo domain.ProjectRepository
}

// NewProjectUsecase NewProjectUsecase
func NewProjectUsecase(a domain.ProjectRepository) domain.ProjectUsecase {
	return &projectUsecase{
		projectRepo: a,
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
	_, err := a.projectRepo.CreateProject(ctx, idUser, url, profilePicture, title, description)
	if err != nil {
		return false, err
	}
	return true, nil
}
