package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/catninpo/gophi"
)

type UserService struct {
	db *sql.DB
}

// NewUserService will create a new instance of a user service with the database
// connection provided. This will then be used for future queries.
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// UserByID will query the users table within the database and return a User
// where their ID matches the given query otherwise if there is any issues
// with the query an error is returned.
func (u *UserService) UserByID(id int) (*gophi.User, error) {
	row := u.db.QueryRow("SELECT id, name, card_id FROM Users WHERE id = $1", id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user gophi.User
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}

	return &user, nil
}
