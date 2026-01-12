package repositories

import "github.com/cropflow/api/internal/domain/entities"

// PersonRepository defines the interface for person data access
type PersonRepository interface {
	Create(person *entities.Person) error
	FindAll() ([]entities.Person, error)
	FindByID(id int64) (*entities.Person, error)
	FindByUsername(username string) (*entities.Person, error)
	Update(person *entities.Person) error
	Delete(id int64) error
}
