package entities

import "time"

// Farm represents a farm entity
type Farm struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	Size      float64   `json:"size" gorm:"not null"`
	Crops     []Crop    `json:"crops,omitempty" gorm:"foreignKey:FarmID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName overrides the default table name
func (Farm) TableName() string {
	return "farms"
}
