package providers

import "database/sql"

type ProviderInterface interface {
	Get(entity string, params map[string]string) *sql.Rows
	Set(entity string, params ...any) bool
}
