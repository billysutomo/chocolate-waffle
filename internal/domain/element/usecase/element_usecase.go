package usecase

import (
	"context"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"go.uber.org/zap"
)

type ElementUsecase struct {
	elementRepo domain.ElementRepository
	logger      *zap.Logger
}

// NewElementUsecase NewElementUsecase
func NewElementUsecase(a domain.ElementRepository, logger *zap.Logger) domain.ElementUsecase {
	return &ElementUsecase{
		elementRepo: a,
		logger:      logger,
	}
}

func (a *ElementUsecase) CreateElement(
	ctx context.Context,
	idProject int,
	ordernum int,
	elementType string,
	elementBody map[string]interface{},
) (bool, error) {
	return true, nil
}
