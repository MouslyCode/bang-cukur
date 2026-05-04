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

	userID := c.MustGet("user_id").(uuid.UUID)
	var transaction transactionModel.Transaction
	var transactionItems []transactionModel.TransactionItem

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		var total int64

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

		transaction = transactionModel.Transaction{
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
		"total":   transaction.Total,
		"paid":    transaction.Paid,
		"change":  transaction.Change,
		"items":   transactionItems,
	})
}

func GetTransactions(c *gin.Context) {
	var transactions []transactionModel.Transaction

	if err := database.DB.Preload("Items").Where("deleted_at IS NULL").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func GetTransactionById(c *gin.Context) {
	idParam := c.Param("id")
	transactionID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID !"})
		return
	}

	var transaction transactionModel.Transaction
	if err := database.DB.Preload("Items").Where("id = ? AND deleted_at IS NULL", transactionID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found!"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func DeleteTransaction(c *gin.Context) {
	idParam := c.Param("id")

	transactionID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Transaction ID!"})
		return
	}

	var transaction transactionModel.Transaction

	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", transactionID).
		First(&transaction).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found!"})
		return
	}

	if err := database.DB.Delete(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully!"})
}

func GetDeletedTransactions(c *gin.Context) {
	var transactions []transactionModel.Transaction
	if err := database.DB.Unscoped().Preload("Items").Where("deleted_at IS NOT NULL").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func RestoreTransaction(c *gin.Context) {
	idParam := c.Param("id")
	transactionID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID!"})
		return
	}

	var transaction transactionModel.Transaction

	if err := database.DB.
		Unscoped().Where("id = ? AND deleted_at IS NOT NULL", transactionID).
		First(&transaction).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found!"})
		return
	}

	if err := database.DB.Unscoped().
		Model(&transaction).Update("deleted_at", nil).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restore transaction!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction restored successfully!"})
}
