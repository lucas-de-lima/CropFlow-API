package dto

import "time"

// CropBodyDTO represents the request body for crop creation
type CropBodyDTO struct {
	Name        string     `json:"name" binding:"required"`
	PlantedArea float64    `json:"plantedArea" binding:"required"`
	FarmID      int64      `json:"farmId,omitempty"` // Removido required pois vem da URL
	PlantedDate *time.Time `json:"plantedDate,omitempty"`
	HarvestDate *time.Time `json:"harvestDate,omitempty"`
}

// CropDTO represents the response for crop data
type CropDTO struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	PlantedArea float64    `json:"plantedArea"`
	FarmID      int64      `json:"farmId"`
	PlantedDate *time.Time `json:"plantedDate,omitempty"`
	HarvestDate *time.Time `json:"harvestDate,omitempty"`
}