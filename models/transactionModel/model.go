package transactionModel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	Total     int64     `gorm:"not null" json:"total"`
	Paid      int64     `gorm:"not null" json:"paid"`
	Change    int64     `gorm:"not null" json:"change"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}

type TransactionItem struct {
	ID            uint      `gorm:"type:char(36);primaryKey;not null" json:"id"`
	TransactionID uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"transaction_id"`
	ItemID        uuid.UUID `gorm:"type:char(36);not null" json:"item_id"`
	ItemName      string    `gorm:"type:varchar(150)" json:"item_name"`
	Price         int64     `gorm:"not null" json:"price"`
	Qty           int       `gorm:"not null" json:"qty"`
	Subtotal      int64     `gorm:"not null" json:"subtotal"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
