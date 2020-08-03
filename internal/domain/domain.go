package domain

import "context"

// AuthUsecase AuthUsecase
type AuthUsecase interface {
	Login(ctx context.Context, username string, password string) (string, error)
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
