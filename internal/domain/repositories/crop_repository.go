package repositories

import "github.com/cropflow/api/internal/domain/entities"

// CropRepository defines the interface for crop data access
type CropRepository interface {
	Create(crop *entities.Crop) error
	FindAll() ([]entities.Crop, error)
	FindByID(id int64) (*entities.Crop, error)
	FindByFarmID(farmID int64) ([]entities.Crop, error)
	Update(crop *entities.Crop) error
	Delete(id int64) error
	AddFertilizer(cropID, fertilizerID int64) error
	FindFertilizersByCropID(cropID int64) ([]entities.Fertilizer, error)
}
