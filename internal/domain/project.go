package domain

import (
	"context"
	"database/sql"
	"time"
)

// ProjectModel ProjectModel
type ProjectModel struct {
	ID             int
	IDUser         int
	URL            string
	ProfilePicture string
	Title          string
	Description    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime
}

// ProjectUsecase ProjectUsecase
type ProjectUsecase interface {
	CreateProject(ctx context.Context, idUser int, url string, profilePicture string, title string, description string) (bool, error)
}

// ProjectRepository ProjectRepository
type ProjectRepository interface {
	CreateProject(ctx context.Context, project ProjectModel) error
}
