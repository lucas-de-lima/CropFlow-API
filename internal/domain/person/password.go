package person

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Password represents a password value object
type Password struct {
	hash string
}

// NewPassword creates a new Password from plain text with validation and hashing
func NewPassword(plainText string) (Password, error) {
	if len(plainText) < 8 {
		return Password{}, errors.New("password must be at least 8 characters")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, err
	}

	return Password{hash: string(hash)}, nil
}

// NewPasswordFromHash creates a Password from an existing hash (for reconstruction from DB)
func NewPasswordFromHash(hash string) Password {
	return Password{hash: hash}
}

// Hash returns the hashed password
func (p Password) Hash() string {
	return p.hash
}

// Compare checks if the plain text password matches the hashed password
func (p Password) Compare(plainText string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(plainText))
}

// IsValid checks if the password is valid (has a hash)
func (p Password) IsValid() bool {
	return p.hash != ""
}
