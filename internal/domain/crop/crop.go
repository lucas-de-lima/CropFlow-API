package crop

import (
	"time"
)

// Crop represents a crop entity
type Crop struct {
	id          int64
	name        string
	plantedArea float64
	farmID      int64
	plantedDate *time.Time
	harvestDate *time.Time
	createdAt   time.Time
	updatedAt   time.Time
}

// NewCrop creates a new Crop with validation (Factory Method)
func NewCrop(name string, plantedArea float64, farmID int64, plantedDate, harvestDate *time.Time) (*Crop, error) {
	if name == "" {
		return nil, ErrInvalidCropName
	}

	if plantedArea <= 0 {
		return nil, ErrInvalidPlantedArea
	}

	if farmID <= 0 {
		return nil, ErrInvalidFarmID
	}

	if plantedDate != nil && harvestDate != nil && harvestDate.Before(*plantedDate) {
		return nil, ErrInvalidHarvestDate
	}

	now := time.Now()
	return &Crop{
		name:        name,
		plantedArea: plantedArea,
		farmID:      farmID,
		plantedDate: plantedDate,
		harvestDate: harvestDate,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

// Restore reconstructs a Crop from persistence (used by repository)
func Restore(id int64, name string, plantedArea float64, farmID int64, plantedDate, harvestDate time.Time, createdAt, updatedAt time.Time) *Crop {
	return &Crop{
		id:          id,
		name:        name,
		plantedArea: plantedArea,
		farmID:      farmID,
		plantedDate: &plantedDate,
		harvestDate: &harvestDate,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

// Getters (encapsulation)
func (c *Crop) ID() int64 {
	return c.id
}

func (c *Crop) Name() string {
	return c.name
}

func (c *Crop) PlantedArea() float64 {
	return c.plantedArea
}

func (c *Crop) FarmID() int64 {
	return c.farmID
}

func (c *Crop) PlantedDate() *time.Time {
	return c.plantedDate
}

func (c *Crop) HarvestDate() *time.Time {
	return c.harvestDate
}

func (c *Crop) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Crop) UpdatedAt() time.Time {
	return c.updatedAt
}

// SetID is used by repository after insertion
func (c *Crop) SetID(id int64) {
	c.id = id
}

// Business Methods

// ChangeName changes the crop name with validation
func (c *Crop) ChangeName(newName string) error {
	if newName == "" {
		return ErrInvalidCropName
	}
	c.name = newName
	c.updatedAt = time.Now()
	return nil
}

// ChangePlantedArea changes the planted area with validation
func (c *Crop) ChangePlantedArea(newArea float64) error {
	if newArea <= 0 {
		return ErrInvalidPlantedArea
	}
	c.plantedArea = newArea
	c.updatedAt = time.Now()
	return nil
}

// SetPlantedDate sets the planting date
func (c *Crop) SetPlantedDate(date *time.Time) {
	c.plantedDate = date
	c.updatedAt = time.Now()
}

// ChangeHarvestDate changes the harvest date with validation
func (c *Crop) ChangeHarvestDate(newDate *time.Time) error {
	if c.plantedDate != nil && newDate != nil && newDate.Before(*c.plantedDate) {
		return ErrInvalidHarvestDate
	}
	c.harvestDate = newDate
	c.updatedAt = time.Now()
	return nil
}

// IsValid checks if the crop is in a valid state
func (c *Crop) IsValid() bool {
	return c.name != "" && c.plantedArea > 0 && c.farmID > 0
}
