package users

// UserServiceRepo defined all methods for service implementation
type UserServiceRepo interface {
	GetUsers() ([]User, error)
	GetUserByEmail(email string) (*Response, error)
}

// UserService object which implementing all methods in UserServiceRepo
type UserService struct {
	uRepo Repository
}

// GetUsers ...
func (u UserService) GetUsers() ([]User, error) {
	return u.uRepo.FindAll()
}

// GetUserByEmail ...
func (u UserService) GetUserByEmail(email string) (*Response, error) {
	// return u.uRepo.ByEmail(email)
	result, err := u.uRepo.ByEmail(email)
	if err != nil {
		return nil, err
	}

	// temp := Response{}.DTO(*result)
	temp := result.DTOResponse()
	return temp, nil
}

// NewService return a UserServiceRepo
func NewService(uRepo Repository) UserServiceRepo {
	return &UserService{
		uRepo: uRepo,
	}
}
