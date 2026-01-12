package mysql

import (
	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/domain/repositories"
	"gorm.io/gorm"
)

type farmRepository struct {
	db *gorm.DB
}

// NewFarmRepository creates a new MySQL farm repository
func NewFarmRepository(db *gorm.DB) repositories.FarmRepository {
	return &farmRepository{db: db}
}

func (r *farmRepository) Create(farm *entities.Farm) error {
	return r.db.Create(farm).Error
}

func (r *farmRepository) FindAll() ([]entities.Farm, error) {
	var farms []entities.Farm
	err := r.db.Find(&farms).Error
	return farms, err
}

func (r *farmRepository) FindByID(id int64) (*entities.Farm, error) {
	var farm entities.Farm
	err := r.db.First(&farm, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &farm, nil
}

func (r *farmRepository) Update(farm *entities.Farm) error {
	return r.db.Save(farm).Error
}

func (r *farmRepository) Delete(id int64) error {
	return r.db.Delete(&entities.Farm{}, id).Error
}
