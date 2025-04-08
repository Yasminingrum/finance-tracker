package handler

import (
	"net/http"
	"time"

	"finance-tracker/internal/domain"
	"finance-tracker/internal/usecase"

	"github.com/gin-gonic/gin"
)

// TransactionHandler handles HTTP requests for transactions
type TransactionHandler struct {
	usecase usecase.TransactionUsecase
}

// NewTransactionHandler creates a new transaction handler
func NewTransactionHandler(u usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{u}
}

// TransactionInput defines the structure for transaction creation requests
type TransactionInput struct {
	Type     string  `json:"type" binding:"required,oneof=income expense"`
	Amount   float64 `json:"amount" binding:"required"`
	Category string  `json:"category"`
	Note     string  `json:"note"`
	Date     string  `json:"date" binding:"required"` // format YYYY-MM-DD
}

// Create handles the creation of a new transaction
func (h *TransactionHandler) Create(c *gin.Context) {
	userID := c.GetUint("user_id") // Retrieved from JWT middleware

	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}

	tx := domain.Transaction{
		Type:     input.Type,
		Amount:   input.Amount,
		Category: input.Category,
		Note:     input.Note,
		Date:     date,
	}

	err = h.usecase.Create(userID, tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "transaction successfully added"})
}

// GetAll retrieves all transactions for the current user
func (h *TransactionHandler) GetAll(c *gin.Context) {
	userID := c.GetUint("user_id")

	txs, err := h.usecase.GetByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": txs})
}
