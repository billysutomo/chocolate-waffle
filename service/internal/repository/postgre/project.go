package postgre

import (
	"context"
	"database/sql"
	"time"

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

func (p *postgreProjectRepository) CreateProject(
	ctx context.Context,
	idUser int,
	url string,
	profilePicture string,
	title string,
	description string,
) (bool, error) {
	var project domain.Project
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
		idUser,
		url,
		profilePicture,
		title,
		description,
		time.Now(),
		nil,
		nil,
	).Scan(&project.ID)

	if err != nil {
		return false, err
	}
	return true, nil
}
