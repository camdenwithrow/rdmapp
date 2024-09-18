package db

import (
	"database/sql"
	"log"
)

type Store interface {
	// GetUser() (*goth.User, error)
	GetUsers()
	// GetUser(id string) (types.User, error)
	// CreateUserSession(*goth.User) (string, error)
}

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) Store {
	return SQLStore{
		db: db,
	}
}

func (store SQLStore) GetUsers() {
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
