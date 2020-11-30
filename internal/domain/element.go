package domain

import (
	"context"
	"database/sql"
	"time"
)

// ElementModel ElementModel
type ElementModel struct {
	ID        int
	IDProject int
	Ordernum  int
	Type      string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// ElementUsecase ElementUsecase
type ElementUsecase interface {
	CreateElement(ctx context.Context, idProject int, ordernum int, elementType string, elementBody string) (bool, error)
}

// ElementRepository ElementRepository
type ElementRepository interface {
	CreateElement(ctx context.Context, element ElementModel) error
}
