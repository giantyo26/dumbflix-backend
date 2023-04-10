package models

type Film struct {
	ID             int      `json:"id" gorm:"primary_key:auto_increment"`
	Title          string   `json:"title" form:"title" gorm:"type: varchar(255)"`
	Film_Thumbnail string   `json:"film_thumbnail" form:"film_thumbnail"`
	Year           int      `json:"year" form:"year" gorm:"type: int"`
	CategoryID     int      `json:"category_Id" form:"category_id"`
	Category       Category `json:"category"`
	Description    string   `json:"description" form:"description" gorm:"type:text"`
}

func (Film) TableName() string {
	return "films"
}
