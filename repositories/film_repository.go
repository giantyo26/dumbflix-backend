package repositories

import (
	"dumbflix-api/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	FindFilms() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	AddFilm(film models.Film) (models.Film, error)
	EditFilm(film models.Film) (models.Film, error)
	DeleteFilm(film models.Film) (models.Film, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFilms() ([]models.Film, error) {
	var film []models.Film
	err := r.db.Preload("Category").Find(&film).Error

	return film, err
}

func (r *repository) GetFilm(ID int) (models.Film, error) {
	var film models.Film
	err := r.db.Preload("Category").First(&film, ID).Error

	return film, err
}

func (r *repository) AddFilm(film models.Film) (models.Film, error) {
	err := r.db.Preload("Category").Create(&film).Error

	return film, err
}

func (r *repository) EditFilm(film models.Film) (models.Film, error) {
	err := r.db.Preload("Category").Save(&film).Error

	return film, err
}

func (r *repository) DeleteFilm(film models.Film) (models.Film, error) {
	err := r.db.Delete(&film).Error

	return film, err
}
