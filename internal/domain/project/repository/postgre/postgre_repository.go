package postgre

import (
	"context"
	"database/sql"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"go.uber.org/zap"
)

type postgreProjectRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewPosgtreProjectRepository NewPosgtreProjectRepository
func NewPosgtreProjectRepository(db *sql.DB, logger *zap.Logger) domain.ProjectRepository {
	return &postgreProjectRepository{db, logger}
}

func (p *postgreProjectRepository) CreateProject(ctx context.Context, project domain.ProjectModel) error {
	// var project domain.Project
	sqlStatement := `INSERT INTO projects (
		id_user, 
		url, 
		profile_picture, 
		title, 
		description, 
		created_at, 
		updated_at, 
		deleted_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	err := p.db.QueryRowContext(
		ctx,
		sqlStatement,
		project.IDUser,
		project.URL,
		project.ProfilePicture,
		project.Title,
		project.Description,
		project.CreatedAt,
		project.UpdatedAt,
		project.DeletedAt,
	).Err()

	if err != nil {
		return err
	}
	return nil
}

func (p *postgreProjectRepository) Get(ctx context.Context, id int) ([]domain.ProjectModel, error) {
	var projects []domain.ProjectModel

	sqlStatement := `SELECT 
		id,
		id_user, 
		url, 
		profile_picture, 
		title,
		description,
		created_at, 
		updated_at, 
		deleted_at
		FROM projects
		WHERE id = $1`

	rows, err := p.db.QueryContext(ctx, sqlStatement, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var project domain.ProjectModel
		if err := rows.Scan(
			&project.ID,
			&project.IDUser,
			&project.URL,
			&project.ProfilePicture,
			&project.Title,
			&project.Description,
			&project.CreatedAt,
			&project.UpdatedAt,
			&project.DeletedAt,
		); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	rerr := rows.Close()
	if rerr != nil {
		return nil, err
	}
	return projects, nil
}
