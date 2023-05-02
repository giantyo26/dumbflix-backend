package transactiondto

import (
	"dumbflix-api/models"
	"time"
)

type TransactionRequest struct {
	StartDate time.Time   `json:"startDate" form:"startDate" `
	DueDate   time.Time   `json:"dueDate" form:"dueDate" `
	UserID    int         `json:"userId" form:"userId" `
	User      models.User `json:"user" form:"user" `
	Status    string      `json:"status" form:"status" gorm:"type: VARCHAR(25)"`
	Price     int         `json:"price"`
}

type TransactionAddRequest struct {
	UserID int    `json:"userId" form:"userId" `
	Price  int    `json:"price"`
	Days   int    `json:"days"`
	Status string `json:"status" form:"status" gorm:"type: VARCHAR(25)" `
}
