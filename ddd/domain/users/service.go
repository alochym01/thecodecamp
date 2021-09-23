package users

// UserServiceRepo defined all methods for service implementation
type UserServiceRepo interface {
	GetUsers() ([]User, error)
	GetUserByEmail(email string) (*User, error)
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
func (u UserService) GetUserByEmail(email string) (*User, error) {
	return u.uRepo.ByEmail(email)
}

// NewUserService return a UserServiceRepo
func NewUserService(uRepo Repository) UserServiceRepo {
	return &UserService{
		uRepo: uRepo,
	}
}
