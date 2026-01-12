package crop

import "errors"

var (
	ErrCropNotFound        = errors.New("crop not found")
	ErrInvalidCropName     = errors.New("invalid crop name: cannot be empty")
	ErrInvalidPlantedArea  = errors.New("invalid planted area: must be greater than zero")
	ErrInvalidFarmID       = errors.New("invalid farm ID")
	ErrInvalidHarvestDate  = errors.New("invalid harvest date: must be after planted date")
	ErrFarmNotFound        = errors.New("farm not found")
	ErrFertilizerNotFound  = errors.New("fertilizer not found")
	ErrDuplicateFertilizer = errors.New("fertilizer already associated with this crop")
)
