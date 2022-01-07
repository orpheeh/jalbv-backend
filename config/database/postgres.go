package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Postgres *sql.DB

func InitPostgres() {
	// Capture connection properties.
	cfg := fmt.Sprintf("postgresql://%v:%v/%v?user=%v&password=%v&sslmode=disable",
		os.Getenv("PG_IP"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DBNAME"),
		os.Getenv("PG_USERNAME"),
		os.Getenv("PG_PASSWORD"))

	// Get a database handle.
	var err error
	Postgres, err = sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := Postgres.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
