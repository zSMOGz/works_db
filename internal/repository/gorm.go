package repository

import (
	"works_db/internal/model"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(name string) (model.User, error) {
	user := model.User{Name: name}
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *GormRepository) Get(id uint) (model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	return user, result.Error
}

func (r *GormRepository) Update(id uint, newName string) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Update("Name", newName).Error
}

func (r *GormRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
