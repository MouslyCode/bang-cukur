package controllers

import (
	"net/http"

	"github.com/MouslyCode/bang-cukur/database"
	"github.com/MouslyCode/bang-cukur/models/productModel"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetProducts(c *gin.Context) {
	var products []productModel.Product
	database.DB.Where("deleted_at IS NULL").Find(&products)
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var input productModel.Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := productModel.Product{
		Name:  input.Name,
		Price: input.Price,
		Img:   input.Img,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    product.ID,
		"Name":  product.Name,
		"Price": product.Price,
		"Img":   product.Img,
	})

}

func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")

	productID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID!"})
		return
	}

	var product productModel.Product
	if err := database.DB.First(&product, "id = ? AND deleted_at IS NULL", productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found!"})
		return
	}
	var input productModel.Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Name = input.Name
	product.Price = input.Price
	product.Img = input.Img

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Update Successfully!",
		"service": product,
	})
}

func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")

	productID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID!"})
		return
	}

	if err := database.DB.Delete(&productModel.Product{}, "id = ?", productID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully!"})
}

func GetDeletedProducts(c *gin.Context) {
	var products []productModel.Product
	database.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&products)
	c.JSON(http.StatusOK, products)
}

func RestoreProduct(c *gin.Context) {
	productID := c.Param("id")
	database.DB.Unscoped().Model(&productModel.Product{}).Where("id = ?", productID).Update("deleted_at", nil)
	c.JSON(http.StatusOK, gin.H{"message": "Product restored successfully"})
}
