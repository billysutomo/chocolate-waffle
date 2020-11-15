package postgre

import (
	"context"
	"database/sql"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
)

type postgreBlockRepository struct {
	db *sql.DB
}

// NewPosgtreBlockRepository NewPosgtreBlockRepository
func NewPosgtreBlockRepository(db *sql.DB) domain.BlockRepository {
	return &postgreBlockRepository{db}
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
