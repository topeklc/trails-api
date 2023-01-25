package database

import (
	"os"
	"fmt"
	"database/sql"
	"io/ioutil"
	"path/filepath"
    _ "github.com/lib/pq"
	"trails-api/models"
)


var DB *sql.DB = ConnectDB()

func getConnectionString() string {
	user := os.Getenv("PGUSER")
	host := os.Getenv("PGHOST")
	pass := os.Getenv("PGPASS")
	port := os.Getenv("PGPORT")
	return fmt.Sprintf("user=%v dbname=postgres  host=%v password=%v port=%v sslmode=disable", user, host, pass, port)
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}
	return db
}

func CreateDB() {
	path, err := os.Getwd()
	if err != nil {
	fmt.Println(err)
	}
	query, err := ioutil.ReadFile(filepath.Join(path, "database/create.sql"))
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}
	if _, err := DB.Exec(string(query)); err != nil {
		fmt.Printf("Error: %v \n", err)
}
}

func CreateUser(user *models.User) {
	query := fmt.Sprintf("INSERT INTO user_account(login, email, password) VALUES (%s, %s, %s)", user.Login, user.Email, user.Password)
	result := DB.QueryRow(query)
	fmt.Printf("Result: %v \n", result)
}