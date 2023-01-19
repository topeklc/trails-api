package main

import (
	"trails-api/database"
	// "github.com/gin-gonic/gin"
)



func main() {
	db := database.ConnectDB()
	database.Migrate(db)
}