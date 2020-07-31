package domain

import "context"

// URLUsecase service
type URLUsecase interface {
	GetURL(ctx context.Context, nama string)
}

// URLRepository repository
type URLRepository interface {
	GetURL(ctx context.Context, nama string)
	// Create(ctx context.Context)
}
