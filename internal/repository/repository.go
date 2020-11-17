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

// UserRepository UserRepository
type UserRepository interface {
	CreateUser(ctx context.Context, user UserModel) error
	GetUserByEmail(ctx context.Context, email string) (UserModel, error)
}

// UserModel UserModel
type UserModel struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
