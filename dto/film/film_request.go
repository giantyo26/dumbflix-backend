package filmdto

type FilmRequest struct {
	Title          string `json:"title" form:"title" validate:"required"`
	Film_Thumbnail string `json:"film_thumbnail" form:"film_thumbnail" validate:"required"`
	Year           int    `json:"year" form:"year" validate:"required"`
	CategoryID     int    `json:"category_id" form:"category" validate:"required"`
	Description    string `json:"description" form:"description" validate:"required"`
}
