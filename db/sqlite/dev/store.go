package dev

import (
	"database/sql"
	"log"

	"github.com/camdenwithrow/rdmapp/db"
)

type DevSQLiteStore struct {
	db *sql.DB
}

func NewDevSQLiteStore() db.Store {
	return DevSQLiteStore{
		db: OpenDevSqliteDB(),
	}
}

func (store DevSQLiteStore) Close() {
	store.db.Close()
}

func (store DevSQLiteStore) GetUsers() {
	rows, err := store.db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}

func (store DevSQLiteStore) GetFeatures() {
	rows, err := store.db.Query("SELECT id, name, description FROM features")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, description string
		err := rows.Scan(&id, &name, &description)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s, Description: %s\n", id, name, description)
	}
}
