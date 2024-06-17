package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func initDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
	}

	script, err := os.ReadFile("./database_setup.sql")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(script))

	if err != nil {
		log.Fatal(err)
	}

	return db
}
