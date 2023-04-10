package repositories

import (
	"dumbflix-api/models"

	"gorm.io/gorm"
)

type EpisodeRepository interface {
	FindEpisodesByFilm(FilmID int) ([]models.Episode, error)
	GetEpisodeByFilm(FilmID int, EpisodeID int) (models.Episode, error)
	GetEpisode(ID int) (models.Episode, error)
	AddEpisode(episode models.Episode) (models.Episode, error)
	EditEpisode(episode models.Episode) (models.Episode, error)
	DeleteEpisode(episode models.Episode, ID int) (models.Episode, error)
}

func RepositoryEpisode(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindEpisodesByFilm(FilmID int) ([]models.Episode, error) {
	var episode []models.Episode
	err := r.db.Preload("Film").Where("film_id = ?", FilmID).Find(&episode).Error

	return episode, err
}

func (r *repository) GetEpisodeByFilm(FilmID int, EpisodeID int) (models.Episode, error) {
	var episode models.Episode
	err := r.db.Preload("Film").Where("film_id = ? AND id = ?", FilmID, EpisodeID).First(&episode).Error

	return episode, err
}

func (r *repository) GetEpisode(ID int) (models.Episode, error) {
	var episode models.Episode

	err := r.db.Preload("Film").First(&episode, ID).Error

	return episode, err
}

func (r *repository) AddEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Preload("Film").Create(&episode).Error

	return episode, err
}

func (r *repository) EditEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Preload("Film").Save(&episode).Error

	return episode, err
}

func (r *repository) DeleteEpisode(episode models.Episode, ID int) (models.Episode, error) {
	err := r.db.Delete(&episode, ID).Scan(&episode).Error

	return episode, err
}
