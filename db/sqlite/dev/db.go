package local

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDevSqliteDB() *sql.DB {
	// Open the database file. If it doesn't exist, create it.
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to SQLite database")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        name TEXT NOT NULL,
                        age INTEGER
                      )`)

	if err != nil {
		fmt.Println("Failed to create User table")
		log.Fatal(err)
	}
	fmt.Println("Created Users Table")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS features (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT,
        status TEXT NOT NULL,
        priority INTEGER NOT NULL,
		target_date TEXT NOT NULL,
		upvotes INTEGER DEFAULT 0,
		date_added TEXT DEFAULT (DATETIME('now')),
        created_at TEXT DEFAULT (DATETIME('now')),
        updated_at TEXT DEFAULT (DATETIME('now'))
    )`)

	if err != nil {
		fmt.Println("Failed to create features table:")
		log.Fatal(err)
	}
	fmt.Println("Created Features Table")

	return db

}
