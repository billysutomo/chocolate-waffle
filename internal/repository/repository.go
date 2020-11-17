package repository

import (
	"context"
	"database/sql"
	"time"
)

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
