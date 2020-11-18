package usecase

import (
	"context"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/domain/block/repository"
	"go.uber.org/zap"
)

type blockUsecase struct {
	blockRepo repository.BlockRepository
	logger    *zap.Logger
}

// NewBlockUsecase NewBlockUsecase
func NewBlockUsecase(a repository.BlockRepository, logger *zap.Logger) domain.BlockUsecase {
	return &blockUsecase{
		blockRepo: a,
		logger:    logger,
	}
}

func (a *blockUsecase) CreateBlock(
	ctx context.Context,
	idProject int,
	ordernum int,
	blockType string,
	blockBody map[string]interface{},
) (bool, error) {
	return true, nil
}
