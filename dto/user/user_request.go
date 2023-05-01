package usersdto

type UserRequest struct {
	Name          string `json:"fullname" gorm:"type: varchar(255)" form:"name"`
	AvatarProfile string `json:"thumbnail" gorm:"type: varchar(255)" form:"thumbnail"`
	Email         string `json:"email" gorm:"type: varchar(255)" form:"email"`
	Password      string `json:"password" gorm:"type :varchar(255)" form:"password"`
	Gender        string `json:"gender" form:"gender"`
	Phone         string `json:"phone" form:"phone"`
	Address       string `json:"address" form:"address"`
	Subscribe     string `json:"subscribe" form:"subscribe"`
}
