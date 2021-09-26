package users

import "github.com/alochym01/thecodecamp_1/errs"

// User domain table
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

// DTOResponse is using for Data Transform Object.
// // DTO is using for Data Transform Object.
// func (res Response) DTO(u User) *Response {
// 	return &Response{
// 		ID:     u.ID,
// 		Email:  u.Email,
// 		Role:   u.Role,
// 		Status: u.statusToText(),
// 	}
// }
func (u User) DTOResponse() *Response {
	return &Response{
		ID:     u.ID,
		Email:  u.Email,
		Role:   u.Role,
		Status: u.statusToText(),
	}
}

// Repository defined all methods for storage implementation
type Repository interface {
	// FindAll() ([]User, error)
	// ByEmail(email string) (*User, error)
	FindAll() ([]User, *errs.AppErrs)
	ByEmail(email string) (*User, *errs.AppErrs)
}

// Request is using for Data Transform Object - DTO.
// As converting data for create user request in form data
// https://github.com/gin-gonic/gin#bind-form-data-request-with-custom-struct
type Request struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Role     string `form:"role" json:"role"`
	Status   string `form:"status" json:"status"`
}

// Response is using for Data Transform Object - DTO. As a response to request
type Response struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

// DTO is using for Data Transform Object.
func (res Response) DTO(u User) *Response {
	return &Response{
		ID:     u.ID,
		Email:  u.Email,
		Role:   u.Role,
		Status: u.statusToText(),
	}
}

// statusToText converting number to string
func (u User) statusToText() string {
	if u.Status == "1" {
		return "Active"
	}
	return "In-Active"
}
