package repositories

import (
	"dumbflix-api/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetTransactionByUser(UserID int) ([]models.Transaction, error)
	AddTransaction(transaction models.Transaction) (models.Transaction, error)
	EditTransaction(transaction models.Transaction) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction

	err := r.db.Preload("User").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) GetTransactionByUser(UserID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Where("user_id = ?", UserID).Find(&transaction).Error

	return transaction, err
}

func (r *repository) AddTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Create(&transaction).Error

	return transaction, err
}

func (r *repository) EditTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Save(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}
