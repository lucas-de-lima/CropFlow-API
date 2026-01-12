package entities

import "time"

// Crop represents a crop entity
type Crop struct {
	ID           int64        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string       `json:"name" gorm:"not null"`
	PlantedArea  float64      `json:"plantedArea" gorm:"column:planted_area;not null"`
	FarmID       int64        `json:"farmId" gorm:"column:farm_id;not null"`
	Farm         *Farm        `json:"farm,omitempty" gorm:"foreignKey:FarmID"`
	PlantedDate  *time.Time   `json:"plantedDate,omitempty" gorm:"column:planting_date"`
	HarvestDate  *time.Time   `json:"harvestDate,omitempty" gorm:"column:harvest_date"`
	Fertilizers  []Fertilizer `json:"fertilizers,omitempty" gorm:"many2many:crop_fertilizer;"`
	CreatedAt    time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName overrides the default table name
func (Crop) TableName() string {
	return "crops"
}
