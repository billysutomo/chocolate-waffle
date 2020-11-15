package postgre

import (
	"context"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type postgreBlockRepository struct {
	Conn   *pgx.Conn
	logger *zap.Logger
}

// NewPosgtreBlockRepository NewPosgtreBlockRepository
func NewPosgtreBlockRepository(Conn *pgx.Conn, logger *zap.Logger) domain.BlockRepository {
	return &postgreBlockRepository{Conn, logger}
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
