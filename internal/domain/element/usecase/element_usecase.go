package usecase

import (
	"context"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"go.uber.org/zap"
)

type elementUsecase struct {
	elementRepo domain.ElementRepository
	logger      *zap.Logger
}

// NewElementUsecase NewElementUsecase
func NewElementUsecase(a domain.ElementRepository, logger *zap.Logger) domain.ElementUsecase {
	return &elementUsecase{
		elementRepo: a,
		logger:      logger,
	}
}

// CreateElement CreateElement
func (a *elementUsecase) CreateElement(
	ctx context.Context,
	idProject int,
	ordernum int,
	elementType string,
	elementBody string,
) (bool, error) {
	element := domain.ElementModel{
		IDProject: idProject,
		Ordernum:  ordernum,
		Type:      elementType,
		Body:      elementBody,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := a.elementRepo.CreateElement(ctx, element)

	if err != nil {
		return false, err
	}
	return true, nil
}
