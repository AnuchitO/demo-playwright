package database

import (
	"database/sql"
	"fmt"

	"demo/config"

	_ "github.com/lib/pq"
)

// PostgresDB returns a new postgres database connection
func NewPostgres(cfg config.Database) (*sql.DB, func()) {
	db, err := sql.Open("postgres", cfg.PostgresURI)
	if err != nil {
		panic(fmt.Errorf("can not connect to postgres: %w", err))
	}

	teardown := func() {
		_ = db.Close()
	}

	return db, teardown
}
