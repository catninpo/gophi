package gophi

// User represents an example user for your usecase.
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// UserService is a way for data layers to interface with your user returning
// it as a result.
type UserService interface {
	UserByID(id int) (*User, error)
}
