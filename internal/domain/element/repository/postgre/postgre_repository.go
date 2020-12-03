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
}

func (p *postgreElementRepository) GetByIDProject(ctx context.Context, idProject int) ([]domain.ElementModel, error) {
	var elements []domain.ElementModel

	sqlStatement := `SELECT 
		id,
		id_project, 
		ordernum, 
		type, 
		body,
		created_at, 
		updated_at, 
		deleted_at
		FROM elements
		WHERE id_project = $1`

	rows, err := p.db.QueryContext(ctx, sqlStatement, idProject)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var element domain.ElementModel
		if err := rows.Scan(
			&element.ID,
			&element.IDProject,
			&element.Ordernum,
			&element.Type,
			&element.Body,
			&element.CreatedAt,
			&element.UpdatedAt,
			&element.DeletedAt,
		); err != nil {
			return nil, err
		}
		elements = append(elements, element)
	}

	rerr := rows.Close()
	if rerr != nil {
		return nil, err
	}
	return elements, nil
}
