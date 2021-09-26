package domain

// UserRepositoryStubs ...
type UserRepositoryStubs struct {
	users []Customer
}

// FindAll ...
func (u UserRepositoryStubs) FindAll() ([]Customer, error) {
	return u.users, nil
}

// NewUserRepositoryStubs ...
func NewUserRepositoryStubs() UserRepositoryStubs {
	temp := []Customer{
		{ID: "1000", Email: "hadn4@fpt.com.vn", Password: "Alochym@123", Role: "Husband", Status: "1"},
		{ID: "1002", Email: "nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	}
	// ID       string
	// Email    string
	// Password string
	// Role     string
	// Status   string
	return UserRepositoryStubs{
		users: temp,
	}
}
