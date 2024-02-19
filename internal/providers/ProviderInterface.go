package providers

import "database/sql"

type ProviderInterface interface {
	Get(table string, params ...any) *sql.Rows
	Set(table string, params ...any) bool
}
