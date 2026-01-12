package usecases

import (
	"errors"

	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/domain/repositories"
)

var (
	ErrFertilizerNotFound = errors.New("fertilizer not found")
)

// FertilizerUseCase handles fertilizer business logic
type FertilizerUseCase struct {
	fertilizerRepo repositories.FertilizerRepository
}

// NewFertilizerUseCase creates a new fertilizer use case
func NewFertilizerUseCase(fertilizerRepo repositories.FertilizerRepository) *FertilizerUseCase {
	return &FertilizerUseCase{
		fertilizerRepo: fertilizerRepo,
	}
}

// CreateFertilizer creates a new fertilizer
func (uc *FertilizerUseCase) CreateFertilizer(fertilizer *entities.Fertilizer) error {
	return uc.fertilizerRepo.Create(fertilizer)
}

// GetAllFertilizers retrieves all fertilizers
func (uc *FertilizerUseCase) GetAllFertilizers() ([]entities.Fertilizer, error) {
	return uc.fertilizerRepo.FindAll()
}

// GetFertilizerByID retrieves a fertilizer by ID
func (uc *FertilizerUseCase) GetFertilizerByID(id int64) (*entities.Fertilizer, error) {
	fertilizer, err := uc.fertilizerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if fertilizer == nil {
		return nil, ErrFertilizerNotFound
	}
	return fertilizer, nil
}

// UpdateFertilizer updates a fertilizer
func (uc *FertilizerUseCase) UpdateFertilizer(fertilizer *entities.Fertilizer) error {
	existing, err := uc.fertilizerRepo.FindByID(fertilizer.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrFertilizerNotFound
	}
	return uc.fertilizerRepo.Update(fertilizer)
}

// DeleteFertilizer deletes a fertilizer
func (uc *FertilizerUseCase) DeleteFertilizer(id int64) error {
	existing, err := uc.fertilizerRepo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrFertilizerNotFound
	}
	return uc.fertilizerRepo.Delete(id)
}
