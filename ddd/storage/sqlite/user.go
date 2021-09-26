package sqlite

import (
	"database/sql"
	"fmt"

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
	sqlstmt := "select * from users"

	rows, err := u.db.Query(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Err", err.Error())
		// return nil, errs.NewServerError(err.Error())
	}

	temp := []users.User{}
	// users := make([]users.User, 0)

	for rows.Next() {
		c := users.User{}
		err := rows.Scan(&c.ID, &c.Email, &c.Password, &c.Role, &c.Status)
		// check err from server DB and Scan function
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Not Found", err.Error())
				// return nil, errs.NewNotFoundError(sql.ErrNoRows.Error())
				return nil, err
			}
			fmt.Println("Server Err", err.Error())
			// return nil, errs.NewServerError("Server Error")
			return nil, err
		}
		temp = append(temp, c)
	}
	return temp, nil
}

// ByEmail ...
func (u Repository) ByEmail(email string) (*users.User, error) {
	sqlstmt := "select * from users where email=?"
	row := u.db.QueryRow(sqlstmt, email)
	c := users.User{}
	err := row.Scan(&c.ID, &c.Email, &c.Password, &c.Role, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Not Found", err.Error())
			return nil, err
			// return nil, errs.NewNotFoundError(sql.ErrNoRows.Error())
		}
		fmt.Println("Server Err", err.Error())
		return nil, err
		// return nil, errs.NewServerError(err.Error())
	}

	return &c, nil
}
