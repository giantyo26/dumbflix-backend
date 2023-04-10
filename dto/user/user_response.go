package usersdto

type DeleteUserResponse struct {
	ID int `json:"id"`
}

type UsersResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password" gorm:"type :varchar(255)"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Subscribe string `json:"subscribe"`
}
