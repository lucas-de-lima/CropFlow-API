package usecases

import (
	"errors"

	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/domain/repositories"
	"github.com/cropflow/api/internal/infrastructure/security"
)

var (
	ErrPersonNotFound      = errors.New("person not found")
	ErrUsernameAlreadyUsed = errors.New("username already in use")
)

// PersonUseCase handles person business logic
type PersonUseCase struct {
	personRepo      repositories.PersonRepository
	passwordService *security.PasswordService
}

// NewPersonUseCase creates a new person use case
func NewPersonUseCase(
	personRepo repositories.PersonRepository,
	passwordService *security.PasswordService,
) *PersonUseCase {
	return &PersonUseCase{
		personRepo:      personRepo,
		passwordService: passwordService,
	}
}

// CreatePerson creates a new person
func (uc *PersonUseCase) CreatePerson(person *entities.Person) error {
	// Check if username already exists
	existing, err := uc.personRepo.FindByUsername(person.Username)
	if err != nil {
		return err
	}
	if existing != nil {
		return ErrUsernameAlreadyUsed
	}

	// Hash password
	hashedPassword, err := uc.passwordService.HashPassword(person.Password)
	if err != nil {
		return err
	}
	person.Password = hashedPassword

	return uc.personRepo.Create(person)
}

// GetAllPersons retrieves all persons
func (uc *PersonUseCase) GetAllPersons() ([]entities.Person, error) {
	return uc.personRepo.FindAll()
}

// GetPersonByID retrieves a person by ID
func (uc *PersonUseCase) GetPersonByID(id int64) (*entities.Person, error) {
	person, err := uc.personRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if person == nil {
		return nil, ErrPersonNotFound
	}
	return person, nil
}

// GetPersonByUsername retrieves a person by username
func (uc *PersonUseCase) GetPersonByUsername(username string) (*entities.Person, error) {
	person, err := uc.personRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if person == nil {
		return nil, ErrPersonNotFound
	}
	return person, nil
}

// UpdatePerson updates a person
func (uc *PersonUseCase) UpdatePerson(person *entities.Person) error {
	existing, err := uc.personRepo.FindByID(person.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrPersonNotFound
	}

	// If password is being updated, hash it
	if person.Password != "" && person.Password != existing.Password {
		hashedPassword, err := uc.passwordService.HashPassword(person.Password)
		if err != nil {
			return err
		}
		person.Password = hashedPassword
	}

	return uc.personRepo.Update(person)
}

// DeletePerson deletes a person
func (uc *PersonUseCase) DeletePerson(id int64) error {
	existing, err := uc.personRepo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrPersonNotFound
	}
	return uc.personRepo.Delete(id)
}
