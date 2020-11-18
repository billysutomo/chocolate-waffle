package domain

import "context"

// BlockUsecase BlockUsecase
type BlockUsecase interface {
	CreateBlock(ctx context.Context, idProject int, ordernum int, blockType string, blockBody map[string]interface{}) (bool, error)
}

// BlockRepository BlockRepository
type BlockRepository interface {
	CreateBlock(ctx context.Context, idProject int, ordernum int, blockType string, blockBody map[string]interface{}) (bool, error)
}
