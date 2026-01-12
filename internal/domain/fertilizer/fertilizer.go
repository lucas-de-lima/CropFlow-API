package fertilizer

import (
	"time"
)

// Fertilizer represents a fertilizer entity
type Fertilizer struct {
	id          int64
	name        string
	brand       string
	composition string
	createdAt   time.Time
	updatedAt   time.Time
}

// NewFertilizer creates a new Fertilizer with validation (Factory Method)
func NewFertilizer(name, brand, composition string) (*Fertilizer, error) {
	if name == "" {
		return nil, ErrInvalidFertilizerName
	}

	if brand == "" {
		return nil, ErrInvalidBrand
	}

	if composition == "" {
		return nil, ErrInvalidComposition
	}

	now := time.Now()
	return &Fertilizer{
		name:        name,
		brand:       brand,
		composition: composition,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

// Restore reconstructs a Fertilizer from persistence (used by repository)
func Restore(id int64, name, brand, composition string, createdAt, updatedAt time.Time) *Fertilizer {
	return &Fertilizer{
		id:          id,
		name:        name,
		brand:       brand,
		composition: composition,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

// Getters (encapsulation)
func (f *Fertilizer) ID() int64 {
	return f.id
}

func (f *Fertilizer) Name() string {
	return f.name
}

func (f *Fertilizer) Brand() string {
	return f.brand
}

func (f *Fertilizer) Composition() string {
	return f.composition
}

func (f *Fertilizer) CreatedAt() time.Time {
	return f.createdAt
}

func (f *Fertilizer) UpdatedAt() time.Time {
	return f.updatedAt
}

// SetID is used by repository after insertion
func (f *Fertilizer) SetID(id int64) {
	f.id = id
}

// Business Methods

// ChangeName changes the fertilizer name with validation
func (f *Fertilizer) ChangeName(newName string) error {
	if newName == "" {
		return ErrInvalidFertilizerName
	}
	f.name = newName
	f.updatedAt = time.Now()
	return nil
}

// ChangeBrand changes the fertilizer brand with validation
func (f *Fertilizer) ChangeBrand(newBrand string) error {
	if newBrand == "" {
		return ErrInvalidBrand
	}
	f.brand = newBrand
	f.updatedAt = time.Now()
	return nil
}

// ChangeComposition changes the fertilizer composition with validation
func (f *Fertilizer) ChangeComposition(newComposition string) error {
	if newComposition == "" {
		return ErrInvalidComposition
	}
	f.composition = newComposition
	f.updatedAt = time.Now()
	return nil
}

// IsValid checks if the fertilizer is in a valid state
func (f *Fertilizer) IsValid() bool {
	return f.name != "" && f.brand != "" && f.composition != ""
}
