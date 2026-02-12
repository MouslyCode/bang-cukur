package controllers

import (
	"net/http"

	"github.com/MouslyCode/bang-cukur/common/constant"
	"github.com/MouslyCode/bang-cukur/database"
	"github.com/MouslyCode/bang-cukur/models/itemModel"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetItems(c *gin.Context) {
	var items []itemModel.Item
	database.DB.Where("deleted_at IS NULL").Find(&items)
	c.JSON(http.StatusOK, items)
}

func GetItemByID(c *gin.Context) {
	idParam := c.Param("id")
	itemID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID!"})
		return
	}

	var item itemModel.Item
	if err := database.DB.Where("id = ? AND deleted_at IS NULL", itemID).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func GetItemByType(c *gin.Context) {
	typeParam := c.Param("type")
	var items []itemModel.Item
	database.DB.Where("type = ? AND deleted_at IS NULL", typeParam).Find(&items)
	c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	var input itemModel.Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Type != constant.Product && input.Type != constant.Service {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type!"})
		return
	}

	if input.Type == constant.Product && input.Stock == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product must have stock"})
		return
	}

	if input.Type == constant.Service {
		input.Stock = nil
	}

	item := itemModel.Item{
		Name:  input.Name,
		Price: input.Price,
		Img:   input.Img,
		Stock: input.Stock,
		Type:  input.Type,
	}

	if err := database.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    item.ID,
		"Name":  item.Name,
		"Price": item.Price,
		"Img":   item.Img,
		"type":  item.Type,
	})

}

func UpdateItem(c *gin.Context) {
	idParam := c.Param("id")

	itemID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID!"})
		return
	}

	var item itemModel.Item
	if err := database.DB.First(&item, "id = ? AND deleted_at IS NULL", itemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}
	var input itemModel.Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Type != constant.Product && input.Type != constant.Service {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type!"})
		return
	}

	if input.Type == constant.Product && input.Stock == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product must have stock"})
		return
	}

	if input.Type == constant.Service {
		input.Stock = nil
	}

	item.Name = input.Name
	item.Price = input.Price
	item.Img = input.Img
	item.Type = input.Type

	if err := database.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item Update Successfully!",
		"item":    item,
	})
}

func DeleteItem(c *gin.Context) {
	idParam := c.Param("id")

	itemID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID!"})
		return
	}

	if err := database.DB.Delete(&itemModel.Item{}, "id = ?", itemID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully!"})
}

func GetDeletedItems(c *gin.Context) {
	var items []itemModel.Item
	database.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&items)
	c.JSON(http.StatusOK, items)
}

func RestoreItem(c *gin.Context) {
	itemID := c.Param("id")
	database.DB.Unscoped().Model(&itemModel.Item{}).Where("id = ?", itemID).Update("deleted_at", nil)
	c.JSON(http.StatusOK, gin.H{"message": "Item restored successfully"})
}
