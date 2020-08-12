package domain

import (
	"context"
	"database/sql"
)

// User User
type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
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
	CreateUser(ctx context.Context, name string, email string, password string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

// URLUsecase service
type URLUsecase interface {
	GetURL(ctx context.Context, nama string)
}

// URLRepository repository
type URLRepository interface {
	GetURL(ctx context.Context, nama string)
	// Create(ctx context.Context)
}
