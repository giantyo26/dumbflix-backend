package filmdto

import "dumbflix-api/models"

type FilmResponse struct {
	Title          string          `json:"title"`
	Film_Thumbnail string          `json:"film_thumbnail"`
	Year           int             `json:"year"`
	CategoryID     int             `json:"category_id"`
	Category       models.Category `json:"category"`
	Description    string          `json:"description"`
}
