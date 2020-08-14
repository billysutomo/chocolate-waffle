package usecase

import (
	"context"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
)

type blockUsecase struct {
	blockRepo domain.BlockRepository
}

// NewBlockUsecase NewBlockUsecase
func NewBlockUsecase(a domain.BlockRepository) domain.BlockUsecase {
	return &blockUsecase{
		blockRepo: a,
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
