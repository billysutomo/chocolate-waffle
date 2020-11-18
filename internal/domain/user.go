package domain

import (
	"context"
	"database/sql"
	"time"
)

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

// UserUsecase AuthUsecase
type UserUsecase interface {
	Login(ctx context.Context, username string, password string) (string, string, error)
	RefreshToken(ctc context.Context, refreshToken string) (string, string, error)
	CreateUser(ctx context.Context, name string, email string, password string) (bool, error)
}

// UserRepository UserRepository
type UserRepository interface {
	CreateUser(ctx context.Context, user UserModel) error
	GetUserByEmail(ctx context.Context, email string) (UserModel, error)
}
