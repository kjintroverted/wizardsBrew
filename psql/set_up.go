package psql

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq" // access to psql driver
)

// NewPostgresConnection attempts to connect to a psql server
// and return a connected DB
func NewPostgresConnection() (db *sql.DB, err error) {
	connection := os.Getenv("PSQL_CONN")
	if connection == "" {
		return nil, errors.New("No connection string found (PSQL_CONN)")
	}

	db, err = sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return
}
