package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	// show table records on browser
	// https://sqliteviewer.app/#/empolyee.db/table/employees/

	db, err := sql.Open("sqlite3", "./empolyee.db")

	if err != nil {
		log.Fatal(err.Error())
	}

	// create empolyee table for one time

	// Create employee table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS employees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		position TEXT NOT NULL,
		salary REAL NOT NULL,
		created_at timestamp NOT NULL,
		updated_at timestamp NOT NULL

);`)

	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10) //set max open connection
	db.SetMaxIdleConns(10) // set max idle connection

	return db
}
