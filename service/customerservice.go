package service

import "github.com/alochym01/thecodecamp/domain"

// CustomerService ...
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

// DefaultCustomerService ...
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer ...
func (cs DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return cs.repo.FindAll()
}

func NewCustomerService(customerRepo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: customerRepo,
	}
}
