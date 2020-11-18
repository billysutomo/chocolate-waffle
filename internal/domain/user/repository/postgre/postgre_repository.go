package postgre

import (
	"context"
	"database/sql"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"go.uber.org/zap"
)

type postgreUserRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewPosgtgreUserRepository NewPosgtgreUserRepository
func NewPosgtgreUserRepository(db *sql.DB, logger *zap.Logger) domain.UserRepository {
	return &postgreUserRepository{db, logger}
}

func (p *postgreUserRepository) CreateUser(ctx context.Context, user domain.UserModel) error {
	sqlStatement := `INSERT INTO users (
		name, 
		email, 
		password, 
		created_at, 
		updated_at, 
		deleted_at
		) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err := p.db.QueryRow(
		sqlStatement,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
	).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgreUserRepository) GetUserByEmail(ctx context.Context, email string) (domain.UserModel, error) {
	var user domain.UserModel
	sqlStatement := `SELECT id, name, email, password, created_at, updated_at, deleted_at FROM users where email = $1`
	err := p.db.QueryRowContext(ctx, sqlStatement, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		return user, err
	}
	return user, nil
}
