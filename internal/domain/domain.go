package domain

import (
	"context"
)

// UserUsecase AuthUsecase
type UserUsecase interface {
	Login(ctx context.Context, username string, password string) (string, string, error)
	RefreshToken(ctc context.Context, refreshToken string) (string, string, error)
	CreateUser(ctx context.Context, name string, email string, password string) (bool, error)
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

// ProjectUsecase ProjectUsecase
type ProjectUsecase interface {
	CreateProject(ctx context.Context, idUser int, url string, profilePicture string, title string, description string) (bool, error)
}

// BlockUsecase BlockUsecase
type BlockUsecase interface {
	CreateBlock(ctx context.Context, idProject int, ordernum int, blockType string, blockBody map[string]interface{}) (bool, error)
}

// BlockRepository BlockRepository
type BlockRepository interface {
	CreateBlock(ctx context.Context, idProject int, ordernum int, blockType string, blockBody map[string]interface{}) (bool, error)
}
