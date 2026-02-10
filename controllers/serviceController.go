package controllers

import (
	"net/http"

	"github.com/MouslyCode/bang-cukur/database"
	"github.com/MouslyCode/bang-cukur/models/serviceModel"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetService(c *gin.Context) {
	var services []serviceModel.Service
	database.DB.Where("deleted_at IS NULL").Find(&services)
	c.JSON(http.StatusOK, services)
}

func CreateService(c *gin.Context) {
	var input serviceModel.Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := serviceModel.Service{
		Name:  input.Name,
		Price: input.Price,
		Img:   input.Img,
	}

	if err := database.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    service.ID,
		"Name":  service.Name,
		"Price": service.Price,
		"Img":   service.Img,
	})

}

func UpdateService(c *gin.Context) {
	idParam := c.Param("id")

	serviceId, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID!"})
		return
	}

	var service serviceModel.Service
	if err := database.DB.First(&service, "id = ? AND deleted_at IS NULL", serviceId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found!"})
		return
	}

	var input serviceModel.Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.Name = input.Name
	service.Price = input.Price
	service.Img = input.Img

	if err := database.DB.Save(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Service Update Successfully!",
		"service": service,
	})
}

func DeleteService(c *gin.Context) {
	idParam := c.Param("id")

	serviceID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID!"})
		return
	}

	if err := database.DB.Delete(&serviceModel.Service{}, "id = ?", serviceID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete service!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully!"})
}

func GetDeletedServices(c *gin.Context) {
	var services []serviceModel.Service
	database.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&services)
	c.JSON(http.StatusOK, services)
}

func RestoreService(c *gin.Context) {
	serviceID := c.Param("id")
	database.DB.Unscoped().Model(&serviceModel.Service{}).Where("id = ?", serviceID).Update("deleted_at", nil)
	c.JSON(http.StatusOK, gin.H{"message": "Service restored successfully"})
}
