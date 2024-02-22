package providers

import "database/sql"

type ProviderInterface interface {
	Get(entity string, params ...any) *sql.Rows
	Set(entity string, params ...any) bool
}
