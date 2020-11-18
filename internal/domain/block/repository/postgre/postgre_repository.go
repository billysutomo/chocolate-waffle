package postgre

import (
	"context"
	"database/sql"

	"github.com/billysutomo/chocolate-waffle/internal/domain/block/repository"
	"go.uber.org/zap"
)

type postgreBlockRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewPosgtreBlockRepository NewPosgtreBlockRepository
func NewPosgtreBlockRepository(db *sql.DB, logger *zap.Logger) repository.BlockRepository {
	return &postgreBlockRepository{db, logger}
}

func (p *postgreBlockRepository) CreateBlock(
	ctx context.Context,
	idProject int,
	ordernum int,
	blockType string,
	blockBody map[string]interface{},
) (bool, error) {

	return true, nil
}
