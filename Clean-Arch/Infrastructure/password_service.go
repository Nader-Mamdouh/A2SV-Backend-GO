package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordService handles password operations.
type PasswordService struct{}

// NewPasswordService creates a new PasswordService instance.
func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

// HashPassword hashes a password.
func (p *PasswordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword verifies a password against a hashed password.
func (p *PasswordService) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
