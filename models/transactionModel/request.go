package transactionModel

import "github.com/google/uuid"

type TransactionItemRequest struct {
	ItemID uuid.UUID `json:"item_id" binding:"required"`
	Qty    *int      `json:"qty" binding:"required,min=1"`
}

type CreateTransactionRequest struct {
	Paid  int64                    `json:"paid" binding:"required"`
	Items []TransactionItemRequest `json:"items" binding:"required,dive"`
}
