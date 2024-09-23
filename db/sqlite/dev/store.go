package dev

import (
	"database/sql"
	"fmt"
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

func (store DevSQLiteStore) GetRoadmap(slug string) (*db.Roadmap, error) {
	var roadmap db.Roadmap
	row := store.db.QueryRow("SELECT * from roadmap where id = ?", slug)
	if err := row.Scan(&roadmap); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No roadmap with slug: %s", slug)
		}
		return nil, fmt.Errorf("Finding roadmap with slug %s failed: %v", slug, err)
	}
	return &roadmap, nil
}

func (store DevSQLiteStore) GetFeatures(roadmapId uint) ([]db.Feature, error) {
	rows, err := store.db.Query("SELECT * FROM features")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	featureArr := []db.Feature{}
	for rows.Next() {
		var feature db.Feature
		err := rows.Scan(&feature)
		if err != nil {
			return nil, err
		}
		featureArr = append(featureArr)
	}
	return featureArr, nil
}
