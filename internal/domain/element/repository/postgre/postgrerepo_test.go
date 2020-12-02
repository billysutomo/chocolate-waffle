package postgre

import (
	"context"
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func NewElementMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreateElement(t *testing.T) {
	db, mock := NewElementMock()
	repo := &postgreElementRepository{db, &zap.Logger{}}
	defer func() {
		repo.db.Close()
	}()

	query := `INSERT INTO elements (
		id_project, 
		ordernum, 
		type, 
		body,
		created_at, 
		updated_at, 
		deleted_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	data := domain.ElementModel{
		IDProject: 1,
		Ordernum:  1,
		Type:      "messenger",
		Body:      "{\"name\":\"john\",\"age\":22,\"class\":\"mca\"}",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.
		NewRows([]string{"id"}).
		AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(
			data.IDProject,
			data.Ordernum,
			data.Type,
			data.Body,
			data.CreatedAt,
			data.UpdatedAt,
			data.DeletedAt,
		).
		WillReturnRows(rows)

	err := repo.CreateElement(context.Background(), data)
	assert.Equal(t, nil, mock.ExpectationsWereMet())
	assert.NoError(t, err)
}
