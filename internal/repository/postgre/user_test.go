package postgre

import (
	"context"
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreateUser(t *testing.T) {
	db, mock := NewMock()
	repo := &postgreUserRepository{db, &zap.Logger{}}
	defer func() {
		repo.db.Close()
	}()

	query := "INSERT INTO users (name, email, password, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	data := UserModel{
		Name:      "Billy",
		Email:     "billysutomo.53@gmail.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.
		NewRows([]string{"id"}).
		AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(
			data.Name,
			data.Email,
			data.Password,
			data.CreatedAt,
			data.UpdatedAt,
			data.DeletedAt,
		).
		WillReturnRows(rows)

	err := repo.CreateUser(context.Background(), data)
	assert.Equal(t, nil, mock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	db, mock := NewMock()
	repo := &postgreUserRepository{db, &zap.Logger{}}
	defer func() {
		repo.db.Close()
	}()

	query := "SELECT id, name, email, password, created_at, updated_at, deleted_at FROM users where email = $1"

	rows := sqlmock.NewRows([]string{
		"id",
		"name",
		"email",
		"password",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		1,
		"Billy",
		"billysutomo.53@gmail.com",
		"passwordd",
		time.Now(),
		time.Now(),
		nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("billysutomo.53@gmail.com").WillReturnRows(rows)
	users, err := repo.GetUserByEmail(context.Background(), "billysutomo.53@gmail.com")
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
}
