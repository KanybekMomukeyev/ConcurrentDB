package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"fmt"
)

func main() {
	db, err := sql.Open("postgres", "dbname=fortest host=localhost sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations/",
		"postgres", driver)

	if err != nil {
		fmt.Println(err)
	}

	m.Up()
}

//migrate create -ext sql -dir migrations noza
//goose postgres "dbname=fortest host=localhost sslmode=disable" status