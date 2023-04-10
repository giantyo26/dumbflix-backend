package episodedto

type EpisodeRequest struct {
	Title             string `json:"title" form:"title" gorm:"type: varchar(255)" validate:"required"`
	Episode_Thumbnail string `json:"image" form:"image" gorm:"type: varchar(255)" validate:"required"`
	Episode_Link      string `json:"video" form:"video" validate:"required"`
	FilmID            int    `json:"film_id" form:"film_id" validate:"required"`
}
