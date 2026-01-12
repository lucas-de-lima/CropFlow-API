package mysql

import (
	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/domain/repositories"
	"gorm.io/gorm"
)

type fertilizerRepository struct {
	db *gorm.DB
}

// NewFertilizerRepository creates a new MySQL fertilizer repository
func NewFertilizerRepository(db *gorm.DB) repositories.FertilizerRepository {
	return &fertilizerRepository{db: db}
}

func (r *fertilizerRepository) Create(fertilizer *entities.Fertilizer) error {
	return r.db.Create(fertilizer).Error
}

func (r *fertilizerRepository) FindAll() ([]entities.Fertilizer, error) {
	var fertilizers []entities.Fertilizer
	err := r.db.Find(&fertilizers).Error
	return fertilizers, err
}

func (r *fertilizerRepository) FindByID(id int64) (*entities.Fertilizer, error) {
	var fertilizer entities.Fertilizer
	err := r.db.First(&fertilizer, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &fertilizer, nil
}

func (r *fertilizerRepository) Update(fertilizer *entities.Fertilizer) error {
	return r.db.Save(fertilizer).Error
}

func (r *fertilizerRepository) Delete(id int64) error {
	return r.db.Delete(&entities.Fertilizer{}, id).Error
}
