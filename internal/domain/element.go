package domain

import "context"

// ElementUsecase ElementUsecase
type ElementUsecase interface {
	CreateElement(ctx context.Context, idProject int, ordernum int, elementType string, elementBody map[string]interface{}) (bool, error)
}

// ElementRepository ElementRepository
type ElementRepository interface {
	CreateElement(ctx context.Context, idProject int, ordernum int, elementType string, elementBody map[string]interface{}) (bool, error)
}
