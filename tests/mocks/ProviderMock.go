package mocks_test

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

type ProviderMock struct{}

func (p *ProviderMock) Get(entity string, params map[string]string) *sql.Rows {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery("^SELECT (.+)").WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).AddRow(1, "test", "test"))

	res, _ := db.Query("SELECT * FROM users WHERE username = $1 AND password = $2", "test", "test")

	return res
}
func (p *ProviderMock) Set(entity string, params ...any) bool {
	return false
}
