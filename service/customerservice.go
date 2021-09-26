package service

import (
	"github.com/alochym01/thecodecamp/domain"
	"github.com/alochym01/thecodecamp/errs"
)

// CustomerService ...
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppErrs)
	GetCustomerByID(string) (*domain.Customer, *errs.AppErrs)
}

// DefaultCustomerService ...
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomer ...
func (cs DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppErrs) {
	return cs.repo.FindAll()
}

// GetCustomer ...
func (cs DefaultCustomerService) GetCustomerByID(id string) (*domain.Customer, *errs.AppErrs) {
	return cs.repo.ByID(id)
}

// NewCustomerService ...
func NewCustomerService(customerRepo domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{
		repo: customerRepo,
	}
}
