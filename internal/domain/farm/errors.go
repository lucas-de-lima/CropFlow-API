package farm

import "errors"

var (
	ErrFarmNotFound        = errors.New("farm not found")
	ErrInvalidFarmName     = errors.New("invalid farm name: cannot be empty")
	ErrInvalidFarmSize     = errors.New("invalid farm size: must be greater than zero")
	ErrFarmCapacityReached = errors.New("farm has reached maximum crop capacity (100)")
	ErrCropNotFound        = errors.New("crop not found in farm")
	ErrDuplicateCrop       = errors.New("crop already exists in farm")
)
