package database

import (
	"os"
	"fmt"
	"database/sql"
	"io/ioutil"
    _ "github.com/lib/pq"
)

func getConnectionString() string {
	user := os.Getenv("PGUSER")
	host := os.Getenv("PGHOST")
	pass := os.Getenv("PGPASS")
	return fmt.Sprintf("user=%v dbname=postgres  host=%v password=%v port=5432 sslmode=disable", user, host, pass)
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}
	return db
}

func Migrate(db *sql.DB) {
	query, err := ioutil.ReadFile("./create.sql")
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}
	if _, err := db.Exec(string(query)); err != nil {
		fmt.Printf("Error: %v \n", err)
}
}