package usecases

import (
	"errors"

	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/domain/repositories"
)

var (
	ErrFarmNotFound = errors.New("farm not found")
)

// FarmUseCase handles farm business logic
type FarmUseCase struct {
	farmRepo repositories.FarmRepository
}

// NewFarmUseCase creates a new farm use case
func NewFarmUseCase(farmRepo repositories.FarmRepository) *FarmUseCase {
	return &FarmUseCase{
		farmRepo: farmRepo,
	}
}

// CreateFarm creates a new farm
func (uc *FarmUseCase) CreateFarm(farm *entities.Farm) error {
	return uc.farmRepo.Create(farm)
}

// GetAllFarms retrieves all farms
func (uc *FarmUseCase) GetAllFarms() ([]entities.Farm, error) {
	return uc.farmRepo.FindAll()
}

// GetFarmByID retrieves a farm by ID
func (uc *FarmUseCase) GetFarmByID(id int64) (*entities.Farm, error) {
	farm, err := uc.farmRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if farm == nil {
		return nil, ErrFarmNotFound
	}
	return farm, nil
}

// UpdateFarm updates a farm
func (uc *FarmUseCase) UpdateFarm(farm *entities.Farm) error {
	existing, err := uc.farmRepo.FindByID(farm.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrFarmNotFound
	}
	return uc.farmRepo.Update(farm)
}

// DeleteFarm deletes a farm
func (uc *FarmUseCase) DeleteFarm(id int64) error {
	existing, err := uc.farmRepo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrFarmNotFound
	}
	return uc.farmRepo.Delete(id)
}
