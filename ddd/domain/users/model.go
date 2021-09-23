package users

// User domain table
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

// Repository defined all methods for storage implementation
type Repository interface {
	FindAll() ([]User, error)
	ByEmail(email string) (*User, error)
}

// Request is using for Data Transform Object - DTO.
// As converting data for create user request data
type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

// Response is using for Data Transform Object - DTO. As a response to request
type Response struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Status string `json:"status"`
}
