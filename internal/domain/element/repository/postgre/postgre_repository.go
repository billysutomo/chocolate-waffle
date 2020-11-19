package postgre

import (
	"context"
	"database/sql"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"go.uber.org/zap"
)

type postgreElementRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewPosgtreElementRepository NewPosgtreElementRepository
func NewPosgtreElementRepository(db *sql.DB, logger *zap.Logger) domain.ElementRepository {
	return &postgreElementRepository{db, logger}
}

func (p *postgreElementRepository) CreateElement(
	ctx context.Context,
	idProject int,
	ordernum int,
	elementType string,
	elementBody map[string]interface{},
) (bool, error) {

	return true, nil
}
