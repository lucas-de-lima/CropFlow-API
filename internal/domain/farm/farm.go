package farm

import (
	"time"
)

// Farm represents a farm aggregate root
type Farm struct {
	id        int64
	name      string
	size      Size
	createdAt time.Time
	updatedAt time.Time
}

// NewFarm creates a new Farm with validation (Factory Method)
func NewFarm(name string, sizeValue float64) (*Farm, error) {
	if name == "" {
		return nil, ErrInvalidFarmName
	}

	size, err := NewSize(sizeValue)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Farm{
		name:      name,
		size:      size,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// Restore reconstructs a Farm from persistence (used by repository)
func Restore(id int64, name string, size Size, createdAt, updatedAt time.Time) *Farm {
	return &Farm{
		id:        id,
		name:      name,
		size:      size,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

// Getters (encapsulation)
func (f *Farm) ID() int64 {
	return f.id
}

func (f *Farm) Name() string {
	return f.name
}

func (f *Farm) Size() Size {
	return f.size
}

func (f *Farm) CreatedAt() time.Time {
	return f.createdAt
}

func (f *Farm) UpdatedAt() time.Time {
	return f.updatedAt
}

// SetID is used by repository after insertion
func (f *Farm) SetID(id int64) {
	f.id = id
}

// Business Methods

// ChangeName changes the farm name with validation
func (f *Farm) ChangeName(newName string) error {
	if newName == "" {
		return ErrInvalidFarmName
	}
	f.name = newName
	f.updatedAt = time.Now()
	return nil
}

// ChangeSize changes the farm size with validation
func (f *Farm) ChangeSize(newSizeValue float64) error {
	newSize, err := NewSize(newSizeValue)
	if err != nil {
		return err
	}
	f.size = newSize
	f.updatedAt = time.Now()
	return nil
}

// IsValid checks if the farm is in a valid state
func (f *Farm) IsValid() bool {
	return f.name != "" && f.size.Value() > 0
}
