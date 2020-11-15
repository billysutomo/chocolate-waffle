package postgre

import (
	"context"
	"database/sql"
	"log"
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

// func TestCreateUser(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &postgreUserRepository{db}
// 	defer func() {
// 		repo.db.Close()
// 	}()

// 	query := "INSERT INTO users (name, email, password, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?)"

// 	myTime := sql.NullTime{Time: time.Now(), Valid: true}
// 	myTimeNil := sql.NullTime{Valid: false}

// 	// mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(query)).
// 		WithArgs(
// 			"Billy",
// 			"billysutomo.53@gmail.com",
// 			"password",
// 			myTime,
// 			myTime,
// 			myTimeNil,
// 		)
// 	// mock.ExpectCommit()

// 	_, err := repo.CreateUser(context.Background(), "Billy", "billysutomo.53@gmail.com", "password", myTime, myTime, myTimeNil)
// 	assert.NoError(t, err)
// }

func TestGetUserByEmail(t *testing.T) {
	db, mock := NewMock()
	repo := &postgreUserRepository{db, &zap.Logger{}}
	defer func() {
		repo.db.Close()
	}()

	query := "SELECT id, name, email, password, created_at, updated_at, deleted_at FROM users where email = ?"

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

	mock.ExpectQuery(query).WithArgs("billysutomo.53@gmail.com").WillReturnRows(rows)
	users, err := repo.GetUserByEmail(context.Background(), "billysutomo.53@gmail.com")
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
}

// 'INSERT INTO users ( name, email, password, created_at, updated_at, deleted_at ) VALUES (?, ?, ?, ?, ?, ?)',
// does not match regex
// 'INSERT INTO users \(name, email, password, created_at, updated_at, deleted_at\) VALUES \(\?, \?, \?, \?, \?, \?\)'
