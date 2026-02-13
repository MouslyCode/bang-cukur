package controllers

import (
	"errors"
	"net/http"

	"github.com/MouslyCode/bang-cukur/common/constant"
	"github.com/MouslyCode/bang-cukur/database"
	"github.com/MouslyCode/bang-cukur/models/itemModel"
	"github.com/MouslyCode/bang-cukur/models/transactionModel"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateTransaction(c *gin.Context) {
	var input transactionModel.CreateTransactionRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("UserID").(uuid.UUID)

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		var total int64
		var transactionItems []transactionModel.TransactionItem

		for _, reqItem := range input.Items {

			var item itemModel.Item
			if err := tx.First(&item, "id = ? AND deleted_at IS NULL", reqItem.ItemID).Error; err != nil {
				return errors.New("Item not found")
			}

			if item.Type == constant.Product {
				if item.Stock == nil || *item.Stock < *reqItem.Qty {
					return errors.New("Insufficient Stock")
				}
			}

			subTotal := item.Price * int64(*reqItem.Qty)
			total += subTotal

			transactionItems = append(transactionItems, transactionModel.TransactionItem{
				ItemID:   item.ID,
				ItemName: item.Name,
				Price:    item.Price,
				Qty:      *reqItem.Qty,
				Subtotal: subTotal,
			})

			if item.Type == constant.Product {
				newStock := *item.Stock - *reqItem.Qty
				if err := tx.Model(&item).Update("stock", newStock).Error; err != nil {
					return err
				}
			}
		}

		if input.Paid < total {
			return errors.New("Insufficient Payment")
		}

		transaction := transactionModel.Transaction{
			UserID: userID,
			Total:  total,
			Paid:   input.Paid,
			Change: input.Paid - total,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		for i := range transactionItems {
			transactionItems[i].TransactionID = transaction.ID
		}

		if err := tx.Create(&transactionItems).Error; err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Transaction Success!",
	})
}
