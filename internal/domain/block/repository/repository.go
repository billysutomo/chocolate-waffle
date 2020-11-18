package repository

import "context"

// BlockRepository BlockRepository
type BlockRepository interface {
	CreateBlock(ctx context.Context, idProject int, ordernum int, blockType string, blockBody map[string]interface{}) (bool, error)
}
