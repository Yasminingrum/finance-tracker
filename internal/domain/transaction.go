// internal/domain/transaction.go
package domain

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Type      string    `gorm:"type:enum('income','expense');not null" json:"type"`
	Amount    float64   `gorm:"not null" json:"amount"`
	Category  string    `gorm:"size:100" json:"category"`
	Note      string    `gorm:"type:text" json:"note"`
	Date      time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `json:"created_at"`

	User User `gorm:"foreignKey:UserID" json:"-"`
}
