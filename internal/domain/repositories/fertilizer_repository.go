package repositories

import "github.com/cropflow/api/internal/domain/entities"

// FertilizerRepository defines the interface for fertilizer data access
type FertilizerRepository interface {
	Create(fertilizer *entities.Fertilizer) error
	FindAll() ([]entities.Fertilizer, error)
	FindByID(id int64) (*entities.Fertilizer, error)
	Update(fertilizer *entities.Fertilizer) error
	Delete(id int64) error
}
