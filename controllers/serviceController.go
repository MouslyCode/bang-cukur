package controllers

import (
	"net/http"

	"github.com/MouslyCode/bang-cukur/database"
	"github.com/MouslyCode/bang-cukur/models/serviceModel"
	"github.com/gin-gonic/gin"
)

func GetService(c *gin.Context) {
	var services []serviceModel.Service
	database.DB.Where("deleted_at IS NULL").Find(&services)
	c.JSON(http.StatusOK, services)
}

// func CreateService(c *gin.Context) {
// 	var input serviceModel.Request
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// }
