package sqlite

import (
	"database/sql"

	"github.com/alochym01/thecodecamp_1/domain/users"
)

// Repository is storage on sqlite and Repository is implemented all method of users.UserRepo
type Repository struct {
	db *sql.DB
}

// NewRepository return a domain.UserRepo
func NewRepository(db *sql.DB) users.Repository {
	return &Repository{
		db: db,
	}
}

// FindAll ...
func (u Repository) FindAll() ([]users.User, error) {
	// result, err := u.db.Query("select * from users")
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(result)
	result1 := []users.User{}
	return result1, nil
}

// ByEmail ...
func (u Repository) ByEmail(email string) (*users.User, error) {
	// user, err := u.db.QueryRow("select * from users where email=?", email)
	return &users.User{}, nil
}
