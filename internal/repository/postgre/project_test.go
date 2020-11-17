package postgre

import (
	"context"
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func NewProjectMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreateProject(t *testing.T) {
	db, mock := NewProjectMock()
	repo := &postgreProjectRepository{db, &zap.Logger{}}
	defer func() {
		repo.db.Close()
	}()

	query := `INSERT INTO projects (
		id_user, 
		url, 
		profile_picture, 
		title, 
		description, 
		created_at, 
		updated_at, 
		deleted_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	data := repository.ProjectModel{
		IDUser:         1,
		URL:            "url",
		ProfilePicture: "ProfilePicture",
		Title:          "title",
		Description:    "description",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	rows := sqlmock.
		NewRows([]string{"id"}).
		AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(
			data.IDUser,
			data.URL,
			data.ProfilePicture,
			data.Title,
			data.Description,
			data.CreatedAt,
			data.UpdatedAt,
			data.DeletedAt,
		).
		WillReturnRows(rows)

	err := repo.CreateProject(context.Background(), data)
	assert.Equal(t, nil, mock.ExpectationsWereMet())
	assert.NoError(t, err)
}
