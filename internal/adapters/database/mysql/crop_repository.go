package mysql

import (
	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/domain/repositories"
	"gorm.io/gorm"
)

type cropRepository struct {
	db *gorm.DB
}

// NewCropRepository creates a new MySQL crop repository
func NewCropRepository(db *gorm.DB) repositories.CropRepository {
	return &cropRepository{db: db}
}

func (r *cropRepository) Create(crop *entities.Crop) error {
	return r.db.Create(crop).Error
}

func (r *cropRepository) FindAll() ([]entities.Crop, error) {
	var crops []entities.Crop
	err := r.db.Find(&crops).Error
	return crops, err
}

func (r *cropRepository) FindByID(id int64) (*entities.Crop, error) {
	var crop entities.Crop
	err := r.db.First(&crop, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &crop, nil
}

func (r *cropRepository) FindByFarmID(farmID int64) ([]entities.Crop, error) {
	var crops []entities.Crop
	err := r.db.Where("farm_id = ?", farmID).Find(&crops).Error
	return crops, err
}

func (r *cropRepository) Update(crop *entities.Crop) error {
	return r.db.Save(crop).Error
}

func (r *cropRepository) Delete(id int64) error {
	return r.db.Delete(&entities.Crop{}, id).Error
}

func (r *cropRepository) AddFertilizer(cropID, fertilizerID int64) error {
	// Load crop and fertilizer to establish the relationship
	crop := &entities.Crop{}
	if err := r.db.First(crop, cropID).Error; err != nil {
		return err
	}

	fertilizer := &entities.Fertilizer{}
	if err := r.db.First(fertilizer, fertilizerID).Error; err != nil {
		return err
	}

	return r.db.Model(crop).Association("Fertilizers").Append(fertilizer)
}

func (r *cropRepository) FindFertilizersByCropID(cropID int64) ([]entities.Fertilizer, error) {
	var crop entities.Crop
	err := r.db.Preload("Fertilizers").First(&crop, cropID).Error
	if err != nil {
		return nil, err
	}
	return crop.Fertilizers, nil
}
