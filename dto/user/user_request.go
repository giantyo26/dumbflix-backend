package usersdto

type UserRequest struct {
	Name      string `json:"fullname" gorm:"type: varchar(255)" form:"name" validate:"required"`
	Email     string `json:"email" gorm:"type: varchar(255)" form:"name" validate:"required"`
	Password  string `json:"password" gorm:"type :varchar(255)" form:"name" validate:"required"`
	Gender    string `json:"gender" form:"name" validate:"required"`
	Phone     string `json:"phone" form:"name" validate:"required"`
	Address   string `json:"address" form:"name" validate:"required"`
	Subscribe string `json:"subscribe" form:"name" validate:"required"`
}
