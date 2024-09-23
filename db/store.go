package db

type Store interface {
	Close()
	// GetUser() (*goth.User, error)
	GetUsers()
	GetRoadmap(slug string) (*Roadmap, error)
	GetFeatures(roadmapId uint) ([]Feature, error)

	// GetUser(id string) (types.User, error)
	// CreateUserSession(*goth.User) (string, error)
}
