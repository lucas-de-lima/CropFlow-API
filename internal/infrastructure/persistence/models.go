package persistence

import "time"

// FarmModel represents the farm database model
type FarmModel struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	Size      float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// TableName overrides the default table name
func (FarmModel) TableName() string {
	return "farms"
}

// CropModel represents the crop database model
type CropModel struct {
	ID          int64      `gorm:"primaryKey;autoIncrement"`
	Name        string     `gorm:"not null"`
	PlantedArea float64    `gorm:"column:planted_area;not null"`
	FarmID      int64      `gorm:"column:farm_id;not null"`
	PlantedDate *time.Time `gorm:"column:planting_date"`
	HarvestDate *time.Time `gorm:"column:harvest_date"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
}

// TableName overrides the default table name
func (CropModel) TableName() string {
	return "crops"
}

// PersonModel represents the person database model
type PersonModel struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// TableName overrides the default table name
func (PersonModel) TableName() string {
	return "person"
}

// FertilizerModel represents the fertilizer database model
type FertilizerModel struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"not null"`
	Brand       string    `gorm:"not null"`
	Composition string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// TableName overrides the default table name
func (FertilizerModel) TableName() string {
	return "fertilizer"
}
