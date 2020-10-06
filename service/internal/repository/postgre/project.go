package postgre

import (
	"context"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/jackc/pgx/v4"
)

type postgreProjectRepository struct {
	Conn *pgx.Conn
}

// NewPosgtreProjectRepository NewPosgtreProjectRepository
func NewPosgtreProjectRepository(Conn *pgx.Conn) domain.ProjectRepository {
	return &postgreProjectRepository{Conn}
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
	err := p.Conn.QueryRow(
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
