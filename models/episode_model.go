package models

type Episode struct {
	ID                int    `json:"id" gorm:"primaryKey"`
	Title             string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Episode_Thumbnail string `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
	Episode_Link      string `json:"episode_link" form:"link"`
	Film              Film   `json:"film" form:"film" gorm:"constraint:OnDelete:CASCADE"`
	FilmID            int    `json:"film_id" form:"film_id"`
}
