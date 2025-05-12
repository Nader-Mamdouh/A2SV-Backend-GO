package usecases

import (
	domain "JWT/Domain"
	repositories "JWT/Repositories"
	"errors"
	"time"
)

// UserUsecase handles user-related business logic.
type UserUsecase struct {
	userRepo repositories.UserRepository
}

// NewUserUsecase creates a new UserUsecase instance.
func NewUserUsecase(userRepo repositories.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

// RegisterUser registers a new user.
func (u *UserUsecase) RegisterUser(user *domain.User) error {
	// Check if user already exists
	existingUser, err := u.userRepo.FindUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("user already exists")
	}

	// Set timestamps
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	// Create user
	return u.userRepo.CreateUser(user)
}

// LoginUser authenticates a user and returns the user if successful.
func (u *UserUsecase) LoginUser(email, password string) (*domain.User, error) {
	user, err := u.userRepo.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// In a real implementation, you would verify the password here.
	// For now, we assume the password is correct.
	return user, nil
}

// GetAllUsers retrieves all users.
func (u *UserUsecase) GetAllUsers() ([]domain.User, error) {
	return u.userRepo.GetAllUsers()
}

// GetUserByID retrieves a user by ID.
func (u *UserUsecase) GetUserByID(id string) (*domain.User, error) {
	return u.userRepo.GetUserByID(id)
}
