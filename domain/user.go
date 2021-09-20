package domain

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
	FindAll() ([]Customer, error)
}
