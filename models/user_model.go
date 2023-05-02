package models

type User struct {
	ID            int    `json:"id"`
	AvatarProfile string `json:"thumbnail" gorm:"type: varchar(255)"`
	Name          string `json:"fullname" gorm:"type: varchar(255)"`
	Email         string `json:"email" gorm:"unique;not null"`
	Password      string `json:"password" gorm:"type :varchar(255)"`
	Gender        string `json:"gender" gorm:"type: varchar(255)"`
	Phone         string `json:"phone" gorm:"type: varchar(255)"`
	Address       string `json:"address" gorm:"type: varchar(255)"`
	Subscribe     string `json:"subscribe"`
	Role          string `json:"roles" gorm:"type: varchar(255)"`
}

type UsersResponse struct {
	ID            int    `json:"id"`
	AvatarProfile string `json:"thumbnail" gorm:"type: varchar(255)"`
	Name          string `json:"fullname"`
	Email         string `json:"email"`
	Password      string `json:"password" gorm:"type :varchar(255)"`
	Gender        string `json:"gender"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	Subscribe     string `json:"subscribe"`
}

func (UsersResponse) TableName() string {
	return "users"
}
