package transactiondto

type TransactionRequest struct {
	StartDate string `json:"startDate" form:"startDate" validate:"required"`
	DueDate   string `json:"dueDate" form:"dueDate" validate:"required"`
	UserID    int    `json:"userId" form:"userId" validate:"required"`
	Status    string `json:"status" form:"status" gorm:"type: VARCHAR(25)" validate:"required"`
}
