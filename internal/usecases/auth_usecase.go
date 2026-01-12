package usecases

import (
	"errors"

	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/domain/repositories"
	"github.com/cropflow/api/internal/infrastructure/security"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// AuthUseCase handles authentication business logic
type AuthUseCase struct {
	personRepo      repositories.PersonRepository
	passwordService *security.PasswordService
	jwtService      *security.JWTService
}

// NewAuthUseCase creates a new auth use case
func NewAuthUseCase(
	personRepo repositories.PersonRepository,
	passwordService *security.PasswordService,
	jwtService *security.JWTService,
) *AuthUseCase {
	return &AuthUseCase{
		personRepo:      personRepo,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

// Login authenticates a user and returns a JWT token
func (uc *AuthUseCase) Login(username, password string) (string, error) {
	// Find user by username
	person, err := uc.personRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if person == nil {
		return "", ErrInvalidCredentials
	}

	// Verify password
	if !uc.passwordService.CheckPassword(person.Password, password) {
		return "", ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := uc.jwtService.GenerateToken(person.Username, string(person.Role))
	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidateToken validates a JWT token and returns the person
func (uc *AuthUseCase) ValidateToken(tokenString string) (*entities.Person, error) {
	// Validate token
	claims, err := uc.jwtService.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	// Find person by username
	person, err := uc.personRepo.FindByUsername(claims.Username)
	if err != nil {
		return nil, err
	}
	if person == nil {
		return nil, ErrPersonNotFound
	}

	return person, nil
}
