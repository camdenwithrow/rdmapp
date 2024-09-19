package db

type Store interface {
	Close()
	// GetUser() (*goth.User, error)
	GetUsers()
	GetFeatures()
	// GetUser(id string) (types.User, error)
	// CreateUserSession(*goth.User) (string, error)
}
