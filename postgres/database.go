package postgres

import (
	"database/sql"
)

// Open wraps the sql.Open and will create a postgres database connection which
// can subsequently be used for any database interactions.
func Open(source string) (*sql.DB, error) {
	return sql.Open("postgres", source)
}
