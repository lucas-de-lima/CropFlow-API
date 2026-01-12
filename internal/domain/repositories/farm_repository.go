package repositories

import "github.com/cropflow/api/internal/domain/entities"

// FarmRepository defines the interface for farm data access
type FarmRepository interface {
	Create(farm *entities.Farm) error
	FindAll() ([]entities.Farm, error)
	FindByID(id int64) (*entities.Farm, error)
	Update(farm *entities.Farm) error
	Delete(id int64) error
}
