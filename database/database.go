package database

import (
	"os"
	"database/sql"
    "github.com/lib/pq"
)

func getConnectionString() string {
	user := os.Getenv("PGUSER")
	host := os.Getenv("PGHOST")
	pass := os.Getenv("PGPASS")
	return fmt.Sprintf("user=%v dbname=postgres  host=%v password=%v port=5432 sslmode=disable", user, host, pass)
}