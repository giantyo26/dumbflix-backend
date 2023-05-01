package transactiondto

import (
	"dumbflix-api/models"
	"time"
)

type TransactionRequest struct {
	StartDate time.Time   `json:"startDate" form:"startDate" validate:"required"`
	DueDate   time.Time   `json:"dueDate" form:"dueDate" validate:"required"`
	UserID    int         `json:"userId" form:"userId" validate:"required"`
	User      models.User `json:"user" form:"user" validate:"required"`
	Status    string      `json:"status" form:"status" gorm:"type: VARCHAR(25)"`
	Price     int         `json:"price"`
}

type TransactionAddRequest struct {
	Price  int    `json:"price"`
	Days   int    `json:"days"`
	Status string `json:"status" form:"status" gorm:"type: VARCHAR(25)" `
}
