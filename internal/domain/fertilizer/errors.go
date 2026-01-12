package fertilizer

import "errors"

var (
	ErrFertilizerNotFound     = errors.New("fertilizer not found")
	ErrInvalidFertilizerName  = errors.New("invalid fertilizer name: cannot be empty")
	ErrInvalidBrand           = errors.New("invalid brand: cannot be empty")
	ErrInvalidComposition     = errors.New("invalid composition: cannot be empty")
)
