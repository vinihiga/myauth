package providers

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

type DatabaseProvider struct {
	db *sql.DB
}

func InitDatabase() (*DatabaseProvider, error) {
	var instance DatabaseProvider = DatabaseProvider{}
	var config string = "host=localhost port=5432 user=admin password=test dbname=test_db sslmode=disable"
	var dbError error = nil

	instance.db, dbError = sql.Open("postgres", config)

	return &instance, dbError
}

func (provider *DatabaseProvider) Get(table string, params map[string]string) *sql.Rows {

	var query string = "SELECT * FROM " + table + " WHERE "
	var args []any = make([]any, 0)
	var paramNum int = 1

	for k, v := range params {
		query += k + " = $" + strconv.Itoa(paramNum)
		paramNum += 1
		args = append(args, v)

		if paramNum <= len(params) {
			query += " AND "
		}
	}

	rows, queryErr := provider.db.Query(query, args...)

	if queryErr != nil {
		return nil
	}

	return rows
}

func (provider *DatabaseProvider) Set(table string, params ...any) bool {
	return false
}
