package crop

import "context"

// Repository defines the interface for crop persistence (Port)
type Repository interface {
	Save(ctx context.Context, crop *Crop) error
	FindByID(ctx context.Context, id int64) (*Crop, error)
	FindAll(ctx context.Context) ([]*Crop, error)
	FindByFarmID(ctx context.Context, farmID int64) ([]*Crop, error)
	Delete(ctx context.Context, id int64) error
	AddFertilizer(ctx context.Context, cropID, fertilizerID int64) error
	FindFertilizersByCropID(ctx context.Context, cropID int64) ([]int64, error)
}
