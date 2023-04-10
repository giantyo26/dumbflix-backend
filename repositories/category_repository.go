package repositories

import (
	"dumbflix-api/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindCategories() ([]models.Category, error)
	GetCategory(ID int) (models.Category, error)
	AddCategory(category models.Category) (models.Category, error)
	EditCategory(category models.Category) (models.Category, error)
	DeleteCategory(category models.Category) (models.Category, error)
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCategories() ([]models.Category, error) {
	var category []models.Category
	err := r.db.Find(&category).Error

	return category, err
}

func (r *repository) GetCategory(ID int) (models.Category, error) {
	var category models.Category

	err := r.db.First(&category, ID).Error

	return category, err
}
func (r *repository) AddCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error

	return category, err
}

func (r *repository) EditCategory(category models.Category) (models.Category, error) {
	err := r.db.Save(&category).Error

	return category, err
}

func (r *repository) DeleteCategory(category models.Category) (models.Category, error) {
	err := r.db.Delete(&category).Error

	return category, err
}
