package domain

import "github.com/alochym01/thecodecamp/errs"

// Customer tables
type Customer struct {
	ID       string
	Email    string
	Password string
	Role     string
	Status   string
}

// CustomerRepository ...
type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppErrs)
	ByID(string) (*Customer, *errs.AppErrs)
}
