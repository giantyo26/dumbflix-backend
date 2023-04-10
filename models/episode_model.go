package models

type Episode struct {
	ID                int    `json:"id" gorm:"primary_key:auto_increment"`
	Title             string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Episode_Thumbnail string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Episode_Link      string `json:"episode_link" form:"link"`
	FilmID            int    `json:"film_id" form:"film_id"`
	Film              Film   `json:"film" form:"film"`
}
