package domain

import (
	"database/sql"
	"fmt"
)

// UserRepositoryDB ...
type UserRepositoryDB struct {
	db *sql.DB
}

// FindAll ...
func (u UserRepositoryDB) FindAll() ([]Customer, error) {
	sqlstmt := "select * from users"

	rows, err := u.db.Query(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	users := []Customer{}
	// users := make([]Customer, 0)

	for rows.Next() {
		c := Customer{}
		err := rows.Scan(&c.ID, &c.Email, &c.Password, &c.Role, &c.Status)
		// check err from server DB and Scan function
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		users = append(users, c)
	}
	return users, nil
}

// NewUserRepositoryDB ...
func NewUserRepositoryDB(dbConn *sql.DB) UserRepositoryDB {
	return UserRepositoryDB{
		db: dbConn,
	}
}
