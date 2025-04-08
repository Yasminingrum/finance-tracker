package usecase

import (
	"finance-tracker/internal/domain"
	"finance-tracker/internal/repository"
)

// TransactionUsecase defines the interface for transaction business logic
type TransactionUsecase interface {
	Create(userID uint, tx domain.Transaction) error
	GetByUser(userID uint) ([]domain.Transaction, error)
}

type transactionUsecase struct {
	repo repository.TransactionRepository
}

// NewTransactionUsecase creates a new transaction usecase
func NewTransactionUsecase(r repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{r}
}

// Create adds a new transaction for a user
func (u *transactionUsecase) Create(userID uint, tx domain.Transaction) error {
	tx.UserID = userID
	return u.repo.Create(&tx)
}

// GetByUser retrieves all transactions for a specific user
func (u *transactionUsecase) GetByUser(userID uint) ([]domain.Transaction, error) {
	return u.repo.GetByUser(userID)
}
