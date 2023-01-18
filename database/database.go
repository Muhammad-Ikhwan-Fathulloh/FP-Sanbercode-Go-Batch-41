package database

import (
	"FP-Sanbercode-Go-Batch-41/database/migrations"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DbConnection *sql.DB
	Err          error
)

func init() {
	Err = godotenv.Load("config/.env")
	if Err != nil {
		fmt.Println("Failed load file environment")
	} else {
		fmt.Println("Success read file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DbConnection, Err = sql.Open("postgres", psqlInfo)
	Err = DbConnection.Ping()
	if Err != nil {
		fmt.Println("Database Connection Failed")
		panic(Err)
	} else {
		fmt.Println("Database Connected")
	}

	statusMigrate := 0
	if statusMigrate == 1 {
		migrations.DbMigrate(DbConnection)

		defer DbConnection.Close()
	}
}
