package entities

import "time"

// Fertilizer represents a fertilizer entity
type Fertilizer struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"not null"`
	Brand       string    `json:"brand" gorm:"not null"`
	Composition string    `json:"composition" gorm:"not null"`
	Crops       []Crop    `json:"crops,omitempty" gorm:"many2many:crop_fertilizer;"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName overrides the default table name
func (Fertilizer) TableName() string {
	return "fertilizer"
}
