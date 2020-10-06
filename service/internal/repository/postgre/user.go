package postgre

import (
	"context"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/jackc/pgx/v4"
)

type postgreUserRepository struct {
	Conn *pgx.Conn
}

// NewPosgtgreUserRepository NewPosgtgreUserRepository
func NewPosgtgreUserRepository(Conn *pgx.Conn) domain.UserRepository {
	return &postgreUserRepository{Conn}
}

func (p *postgreUserRepository) CreateUser(ctx context.Context, name string, email string, password string) (bool, error) {
	var user domain.User
	sqlStatement := `INSERT INTO users (
		name, 
		email, 
		password, 
		created_at, 
		updated_at, 
		deleted_at
	) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := p.Conn.QueryRow(ctx, sqlStatement, name, email, password, time.Now(), nil, nil).Scan(&user.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *postgreUserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	sqlStatement := `SELECT id, name, email, password, created_at, updated_at, deleted_at FROM users where email = $1`
	err := p.Conn.QueryRow(ctx, sqlStatement, email).Scan(
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
