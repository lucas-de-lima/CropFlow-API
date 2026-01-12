package mysql

import (
	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/domain/repositories"
	"gorm.io/gorm"
)

type personRepository struct {
	db *gorm.DB
}

// NewPersonRepository creates a new MySQL person repository
func NewPersonRepository(db *gorm.DB) repositories.PersonRepository {
	return &personRepository{db: db}
}

func (r *personRepository) Create(person *entities.Person) error {
	return r.db.Create(person).Error
}

func (r *personRepository) FindAll() ([]entities.Person, error) {
	var persons []entities.Person
	err := r.db.Find(&persons).Error
	return persons, err
}

func (r *personRepository) FindByID(id int64) (*entities.Person, error) {
	var person entities.Person
	err := r.db.First(&person, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &person, nil
}

func (r *personRepository) FindByUsername(username string) (*entities.Person, error) {
	var person entities.Person
	err := r.db.Where("username = ?", username).First(&person).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &person, nil
}

func (r *personRepository) Update(person *entities.Person) error {
	return r.db.Save(person).Error
}

func (r *personRepository) Delete(id int64) error {
	return r.db.Delete(&entities.Person{}, id).Error
}
