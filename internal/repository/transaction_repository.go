package repository

import (
	"finance-tracker/internal/domain"

	"gorm.io/gorm"
)

// TransactionRepository defines the interface for transaction data access
type TransactionRepository interface {
	Create(tx *domain.Transaction) error
	GetByUser(userID uint) ([]domain.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new transaction repository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

// Create adds a new transaction to the database
func (r *transactionRepository) Create(tx *domain.Transaction) error {
	return r.db.Create(tx).Error
}

// GetByUser retrieves all transactions for a specific user
func (r *transactionRepository) GetByUser(userID uint) ([]domain.Transaction, error) {
	var txs []domain.Transaction
	err := r.db.Where("user_id = ?", userID).Order("date desc").Find(&txs).Error
	return txs, err
}
