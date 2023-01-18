package main

import (
	"FP-Sanbercode-Go-Batch-41/database"
	"FP-Sanbercode-Go-Batch-41/router"
	"database/sql"
	"os"
)

var (
	DB *sql.DB
)

func main() {
	database.DbMigrate(DB)

	defer DB.Close()

	router.StartServer().Run(":" + os.Getenv("PORT"))
}
