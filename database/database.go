package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
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

	// defer DbConnection.Close()
}

func DbMigrate(dbParam *sql.DB) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./migrations"),
	}

	n, errors := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errors != nil {
		panic(errors)
	}

	DbConnection = dbParam

	fmt.Println("Applied ", n, " migrations!")
}
