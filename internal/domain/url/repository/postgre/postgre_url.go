package postgre

import (
	"context"
	"log"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/jackc/pgx/v4"
)

type postgreURLRepository struct {
	Conn *pgx.Conn
}

// NewPostgreURLRepository aa
func NewPostgreURLRepository(Conn *pgx.Conn) domain.URLRepository {
	return &postgreURLRepository{Conn}
}

func (p *postgreURLRepository) GetURL(ctx context.Context, nama string) {
	log.Printf(nama)
}
