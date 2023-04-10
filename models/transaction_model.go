package models

type Transaction struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	StartDate string `json:"startDate"`
	DueDate   string `json:"dueDate"`
	UserID    int    `json:"userId"`
	User      User   `json:"user"`
	Status    string `json:"status"  gorm:"type:varchar(25)"`
}

func (Transaction) TableName() string {
	return "transactions"
}
