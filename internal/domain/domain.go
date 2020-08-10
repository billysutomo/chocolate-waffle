package domain

import "context"

// UserUsecase AuthUsecase
type UserUsecase interface {
	Login(ctx context.Context, username string, password string) (string, string, error)
	RefreshToken(ctc context.Context, refreshToken string) (string, string, error)
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
