package migrations

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateTables(db *sql.DB) {
	users := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT UNIQUE NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`

	owners := `CREATE TABLE IF NOT EXISTS owners (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT UNIQUE NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users (id)
	)`

	roadmapContributors := `CREATE TABLE IF NOT EXISTS roadmap_contributors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    roadmap_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (roadmap_id) REFERENCES roadmaps (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
	)`

	roadmaps := `CREATE TABLE IF NOT EXISTS roadmaps (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
	owner_id INTEGER NOT NULL,
    slug TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
	logo TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES owners (id)
	)`

	features := `CREATE TABLE IF NOT EXISTS features (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    roadmap_id INTEGER NOT NULL,
    name TEXT NOT NULL,
	description TEXT,
    status TEXT NOT NULL,
	priority INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (roadmap_id) REFERENCES roadmaps (id)
	)`

	votes := `CREATE TABLE IF NOT EXISTS votes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    feature_id INTEGER NOT NULL,
    roadmap_contributor_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (feature_id) REFERENCES features (id),
    FOREIGN KEY (roadmap_contributor_id) REFERENCES roadmap_contributors (id),
    UNIQUE (feature_id, roadmap_contributor_id)
	)`

	comments := `CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    feature_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    comment_text TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (feature_id) REFERENCES features (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
	)`

	cmds := map[string]string{
		"users":               users,
		"owners":              owners,
		"roadmapContributors": roadmapContributors,
		"roadmaps":            roadmaps,
		"features":            features,
		"votes":               votes,
		"comments":            comments,
	}

	for k, v := range cmds {
		_, err := db.Exec(v)
		if err != nil {
			fmt.Printf("Failed to create %s table\n", k)
			log.Fatal(err)
		}
		fmt.Printf("Created %s Table\n", k)
	}
}

func InsertMockData(db *sql.DB) {
	mockUsers := []struct {
		email string
	}{
		{"user1@example.com"}, {"user1@example.com"}, {"admin@example.com"},
	}
	mockRoadmaps := []struct {
		ownerId uint
		slug    string
		title   string
		logo    string
	}{{1, "roadmap1", "roadmap1", "https://placehold.co/60x40"}}
	mockFeatures := []struct {
		name        string
		description string
		status      string
		priority    int
	}{
		{"User Profiles", "Allow users to create and edit their profiles", "planned", 1},
		{"Dark Mode", "Add a dark mode theme for the application", "in progress", 2},
		{"Analytics Dashboard", "Provide users with an analytics dashboard", "planned", 3},
		{"Email Notifications", "Send email notifications for important events", "completed", 2},
		{"Multi-language Support", "Support multiple languages in the app", "planned", 4},
	}

	// SQL statement to insert a feature
	insertUserSQL := `INSERT INTO users (email) VALUES (?);`
	insertRoadmapSQL := `INSERT INTO roadmaps (owner_id, slug, title, logo) VALUES (?, ?, ?, ?);`
	insertFeatureSQL := `INSERT INTO features (roadmap_id, name, description, status, priority) VALUES (?, ?, ?, ?, ?);`

	for _, user := range mockUsers {
		_, err := db.Exec(insertUserSQL, user.email)
		if err != nil {
			log.Printf("Error inserting user '%s': %v", user.email, err)
		} else {
			log.Printf("Inserted user: %s", user.email)
		}
	}
	var lastRoadmapId int64
	for _, roadmap := range mockRoadmaps {
		log.Println("roadmaps")
		res, err := db.Exec(insertRoadmapSQL, roadmap.ownerId, roadmap.slug, roadmap.title, roadmap.logo)
		if err != nil {
			log.Printf("Error inserting roadmap '%s': %v", roadmap.title, err)
		} else {
			log.Printf("Inserted roadmap: %s", roadmap.title)
		}
		lastRoadmapId, err = res.LastInsertId()
		if err != nil {
			log.Printf("Error getting last insert id: %v", err)
		}
		log.Printf("Got roadmap id: %d", lastRoadmapId)
	}
	for _, feature := range mockFeatures {
		_, err := db.Exec(insertFeatureSQL, lastRoadmapId, feature.name, feature.description, feature.status, feature.priority)
		if err != nil {
			log.Printf("Error inserting feature '%s': %v", feature.name, err)
		} else {
			log.Printf("Inserted feature: %s", feature.name)
		}
	}
}
