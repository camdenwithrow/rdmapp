package db

import "time"

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Owner struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoadmapContributor struct {
	ID        uint      `json:"id"`
	RoadmapID uint      `json:"roadmap_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Roadmap struct {
	ID        uint      `json:"id"`
	OwnerId   uint      `json:"owner_id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Feature struct {
	ID          uint      `json:"id"`
	RoadmapID   uint      `json:"roadmap_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    int       `json:"priority"`
	DateAdded   time.Time `json:"date_added"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Vote struct {
	ID                   uint      `json:"id"`
	FeatureID            uint      `json:"feature_id"`
	RoadmapContributorID uint      `json:"roadmap_contributor_id"`
	CreatedAt            time.Time `json:"created_at"`
}

type Comment struct {
	ID          uint      `json:"id"`
	FeatureID   uint      `json:"feature_id"`
	UserID      uint      `json:"user_id"`
	CommentText string    `json:"comment_text"`
	CreatedAt   time.Time `json:"created_at"`
}
