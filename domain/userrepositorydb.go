package domain

import (
	"database/sql"
	"fmt"

	"github.com/alochym01/thecodecamp/errs"
)

// UserRepositoryDB ...
type UserRepositoryDB struct {
	db *sql.DB
}

// FindAll ...
func (u UserRepositoryDB) FindAll() ([]Customer, *errs.AppErrs) {
	sqlstmt := "select * from users"

	rows, err := u.db.Query(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Err", err.Error())
		return nil, errs.NewServerError(err.Error())
	}

	users := []Customer{}
	// users := make([]Customer, 0)

	for rows.Next() {
		c := Customer{}
		err := rows.Scan(&c.ID, &c.Email, &c.Password, &c.Role, &c.Status)
		// check err from server DB and Scan function
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError(sql.ErrNoRows.Error())
			}
			return nil, errs.NewServerError("Server Error")
		}
		users = append(users, c)
	}
	return users, nil
}

// ByID ...
func (u UserRepositoryDB) ByID(id string) (*Customer, *errs.AppErrs) {
	sqlstmt := "select * from users where id=?"

	row := u.db.QueryRow(sqlstmt, id)
	c := Customer{}
	err := row.Scan(&c.ID, &c.Email, &c.Password, &c.Role, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError(sql.ErrNoRows.Error())
		}
		return nil, errs.NewServerError(err.Error())
	}

	return &c, nil
}

// NewUserRepositoryDB ...
func NewUserRepositoryDB(dbConn *sql.DB) CustomerRepository {
	return UserRepositoryDB{
		db: dbConn,
	}
}
