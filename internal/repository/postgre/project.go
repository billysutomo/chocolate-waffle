package postgre

import (
	"context"
	"database/sql"

	"github.com/billysutomo/chocolate-waffle/internal/repository"
	"go.uber.org/zap"
)

type postgreProjectRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewPosgtreProjectRepository NewPosgtreProjectRepository
func NewPosgtreProjectRepository(db *sql.DB, logger *zap.Logger) repository.ProjectRepository {
	return &postgreProjectRepository{db, logger}
}

func (p *postgreProjectRepository) CreateProject(ctx context.Context, project repository.ProjectModel) error {
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
	).Scan(&project.ID)

	if err != nil {
		return err
	}
	return nil
}
