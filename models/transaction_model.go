package models

import "time"

type Transaction struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time `json:"startDate"`
	DueDate   time.Time `json:"dueDate"`
	UserID    int       `json:"userId"`
	User      User      `json:"user" gorm:"constraint:OnDelete:SET NULL;"`
	Price     int       `json:"price"`
	Status    string    `json:"status"  gorm:"type:varchar(25)"`
}

func (Transaction) TableName() string {
	return "transactions"
}
