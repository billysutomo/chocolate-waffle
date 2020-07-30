package postgre

import (
	"database/sql"
	"github.com/billysutomo/chocolate-waffle/internal/domain"
)

type postgreURLRepository struct {
	Conn *sql.DB
}

// NewPostgreURLRepository aa
func NewPostgreURLRepository(Conn *sql.DB) domain.URLRepository {
	return &postgreURLRepository{Conn}
}
