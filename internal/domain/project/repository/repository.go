package repository

import (
	"context"
	"database/sql"
	"time"
)

// ProjectRepository ProjectRepository
type ProjectRepository interface {
	CreateProject(ctx context.Context, project ProjectModel) error
}

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
