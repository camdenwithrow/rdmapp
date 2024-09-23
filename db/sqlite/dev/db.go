package dev

import (
	"database/sql"
	"log"

	"github.com/camdenwithrow/rdmapp/db/sqlite/dev/migrations"
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

	migrations.CreateTables(db)

	migrations.InsertMockData(db)

	return db
}

// func insertMockFeatures(db *sql.DB) {
// 	mockFeatures := []struct {
// 		name        string
// 		description string
// 		status      string
// 		priority    int
// 	}{
// 		{"User Profiles", "Allow users to create and edit their profiles", "planned", 1},
// 		{"Dark Mode", "Add a dark mode theme for the application", "in progress", 2},
// 		{"Analytics Dashboard", "Provide users with an analytics dashboard", "planned", 3},
// 		{"Email Notifications", "Send email notifications for important events", "completed", 2},
// 		{"Multi-language Support", "Support multiple languages in the app", "planned", 4},
// 	}
//
// 	// SQL statement to insert a feature
// 	insertFeatureSQL := `INSERT INTO features (name, description, status, priority) VALUES (?, ?, ?, ?);`
//
// 	for _, feature := range mockFeatures {
// 		_, err := db.Exec(insertFeatureSQL, feature.name, feature.description, feature.status, feature.priority)
// 		if err != nil {
// 			log.Printf("Error inserting feature '%s': %v", feature.name, err)
// 		} else {
// 			log.Printf("Inserted feature: %s", feature.name)
// 		}
// 	}
// }
