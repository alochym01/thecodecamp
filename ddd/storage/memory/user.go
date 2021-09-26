package memory

import (
	"github.com/alochym01/thecodecamp_1/domain/users"
	"github.com/alochym01/thecodecamp_1/errs"
)

var temp = []users.User{
	{ID: "1000", Email: "hadn4@fpt.com.vn", Password: "Alochym@123", Role: "Husband", Status: "1"},
	{ID: "1002", Email: "nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
}

// Repository is storage on memory and Repository is implemented all method of users.UserRepo
type Repository struct {
	users []users.User
}

// NewRepository return a users.Repository
func NewRepository() users.Repository {
	return &Repository{
		users: temp,
	}
}

// FindAll ...
func (u Repository) FindAll() ([]users.User, *errs.AppErrs) {
	// result, err := u.db.Query("select * from users")
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(result)

	return u.users, nil
}

// ByEmail ...
func (u Repository) ByEmail(email string) (*users.User, *errs.AppErrs) {
	// user, err := u.db.QueryRow("select * from users where email=?", email)
	for i := range temp {
		if temp[i].Email == email {
			return &temp[i], nil
		}
	}
	return nil, errs.NewNotFoundError("Not Found")
}
