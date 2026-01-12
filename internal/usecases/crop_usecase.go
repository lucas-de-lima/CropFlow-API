package usecases

import (
	"errors"

	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/domain/repositories"
)

var (
	ErrCropNotFound = errors.New("crop not found")
)

// CropUseCase handles crop business logic
type CropUseCase struct {
	cropRepo       repositories.CropRepository
	farmRepo       repositories.FarmRepository
	fertilizerRepo repositories.FertilizerRepository
}

// NewCropUseCase creates a new crop use case
func NewCropUseCase(
	cropRepo repositories.CropRepository,
	farmRepo repositories.FarmRepository,
	fertilizerRepo repositories.FertilizerRepository,
) *CropUseCase {
	return &CropUseCase{
		cropRepo:       cropRepo,
		farmRepo:       farmRepo,
		fertilizerRepo: fertilizerRepo,
	}
}

// CreateCrop creates a new crop
func (uc *CropUseCase) CreateCrop(crop *entities.Crop) error {
	// Validate that farm exists
	farm, err := uc.farmRepo.FindByID(crop.FarmID)
	if err != nil {
		return err
	}
	if farm == nil {
		return ErrFarmNotFound
	}
	return uc.cropRepo.Create(crop)
}

// GetAllCrops retrieves all crops
func (uc *CropUseCase) GetAllCrops() ([]entities.Crop, error) {
	return uc.cropRepo.FindAll()
}

// GetCropByID retrieves a crop by ID
func (uc *CropUseCase) GetCropByID(id int64) (*entities.Crop, error) {
	crop, err := uc.cropRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if crop == nil {
		return nil, ErrCropNotFound
	}
	return crop, nil
}

// GetCropsByFarmID retrieves all crops for a specific farm
func (uc *CropUseCase) GetCropsByFarmID(farmID int64) ([]entities.Crop, error) {
	// Validate that farm exists
	farm, err := uc.farmRepo.FindByID(farmID)
	if err != nil {
		return nil, err
	}
	if farm == nil {
		return nil, ErrFarmNotFound
	}
	return uc.cropRepo.FindByFarmID(farmID)
}

// UpdateCrop updates a crop
func (uc *CropUseCase) UpdateCrop(crop *entities.Crop) error {
	existing, err := uc.cropRepo.FindByID(crop.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrCropNotFound
	}
	return uc.cropRepo.Update(crop)
}

// DeleteCrop deletes a crop
func (uc *CropUseCase) DeleteCrop(id int64) error {
	existing, err := uc.cropRepo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrCropNotFound
	}
	return uc.cropRepo.Delete(id)
}

// AddFertilizerToCrop associates a fertilizer with a crop
func (uc *CropUseCase) AddFertilizerToCrop(cropID, fertilizerID int64) error {
	// Validate crop exists
	crop, err := uc.cropRepo.FindByID(cropID)
	if err != nil {
		return err
	}
	if crop == nil {
		return ErrCropNotFound
	}

	// Validate fertilizer exists
	fertilizer, err := uc.fertilizerRepo.FindByID(fertilizerID)
	if err != nil {
		return err
	}
	if fertilizer == nil {
		return errors.New("fertilizer not found")
	}

	return uc.cropRepo.AddFertilizer(cropID, fertilizerID)
}

// GetFertilizersByCropID retrieves all fertilizers for a specific crop
func (uc *CropUseCase) GetFertilizersByCropID(cropID int64) ([]entities.Fertilizer, error) {
	// Validate crop exists
	crop, err := uc.cropRepo.FindByID(cropID)
	if err != nil {
		return nil, err
	}
	if crop == nil {
		return nil, ErrCropNotFound
	}

	return uc.cropRepo.FindFertilizersByCropID(cropID)
}
