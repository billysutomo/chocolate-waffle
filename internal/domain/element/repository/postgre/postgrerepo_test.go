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

func TestGetByIDProject(t *testing.T) {
	db, mock := NewElementMock()
	repo := &postgreElementRepository{db, &zap.Logger{}}
	defer func() {
		repo.db.Close()
	}()

	query := `SELECT 
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

	rows := sqlmock.NewRows([]string{
		"id",
		"id_project",
		"ordernum",
		"type",
		"body",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		1,
		1,
		1,
		"messenger",
		"body",
		time.Now(),
		time.Now(),
		nil,
	).AddRow(
		2,
		1,
		2,
		"messenger",
		"body",
		time.Now(),
		time.Now(),
		nil)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)
	elements, err := repo.GetByIDProject(context.Background(), 1)

	assert.NotEmpty(t, elements)
	assert.NoError(t, err)
}
