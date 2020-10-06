package postgre

import (
	"context"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/jackc/pgx/v4"
)

type postgreBlockRepository struct {
	Conn *pgx.Conn
}

// NewPosgtreBlockRepository NewPosgtreBlockRepository
func NewPosgtreBlockRepository(Conn *pgx.Conn) domain.BlockRepository {
	return &postgreBlockRepository{Conn}
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
