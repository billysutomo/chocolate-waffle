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
	element domain.ElementModel,
) error {

	sqlStatement := `INSERT INTO elements (
		id_project, 
		ordernum, 
		type, 
		body,
		created_at, 
		updated_at, 
		deleted_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err := p.db.QueryRow(
		sqlStatement,
		element.IDProject,
		element.Ordernum,
		element.Type,
		element.Body,
		element.CreatedAt,
		element.UpdatedAt,
		element.DeletedAt,
	).Scan(&element.ID)
	if err != nil {
		return err
	}
	return nil

	return nil
}
