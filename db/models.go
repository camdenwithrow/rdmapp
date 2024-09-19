package db

type Feature struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    int    `json:"priority"`
	TargetDate  string `json:"target_date"`
	UpVotes     int    `json:"upvotes"`
	DateAdded   string `json:"date_added"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
