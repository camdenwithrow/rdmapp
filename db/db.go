package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenSqliteDB() *sql.DB {
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
		log.Fatal(err)
	}

	return db

}

func insertData(db *sql.DB, age int) {
	// Insert some data
	stmt, err := db.Prepare("INSERT INTO users(name, age) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("Alice", age)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data inserted successfully")
}
