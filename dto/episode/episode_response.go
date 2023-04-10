package episodedto

import "dumbflix-api/models"

type EpisodeResponse struct {
	ID        int         `json:"id,"`
	Title     string      `json:"title"`
	Thumbnail string      `json:"thumbnail"`
	VideoLink string      `json:"video_link"`
	FilmID    int         `json:"film_id"`
	Film      models.Film `json:"film"`
}
