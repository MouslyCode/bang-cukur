package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/MouslyCode/bang-cukur/common/constant"
	"github.com/MouslyCode/bang-cukur/common/helper"
	"github.com/MouslyCode/bang-cukur/database"
	"github.com/MouslyCode/bang-cukur/models/userModel"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) {
	var users []userModel.User
	database.DB.Where("deleted_at IS NULL").Find(&users)
	c.JSON(http.StatusOK, users)
}

func Login(c *gin.Context) {
	var input userModel.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := strings.ToLower(input.Email)

	var user userModel.User
	err := database.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email!"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if !helper.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := helper.GenerateJWT(user.ID, user.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Login sucessfully",
		"token":   token,
	})

}

func CreateUser(c *gin.Context) {
	var input userModel.CreateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.RoleID != constant.RoleOwnerID && input.RoleID != constant.RoleCashierID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	email := strings.ToLower(input.Email)
	var existing userModel.User
	if err := database.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already existed!"})
		return
	}

	user := userModel.User{
		Name:     input.Name,
		Email:    email,
		Password: input.Password,
		RoleID:   input.RoleID,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role": map[uint]string{
			constant.RoleOwnerID:   "owner",
			constant.RoleCashierID: "cashier",
		}[user.RoleID],
	})

}

func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	userID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID!"})
		return
	}

	var user userModel.User
	if err := database.DB.First(&user, "id = ? AND deleted_at IS NULL", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input userModel.UpdateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.RoleID != constant.RoleOwnerID && input.RoleID != constant.RoleCashierID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	user.Name = input.Name
	user.Email = strings.ToLower(input.Email)
	user.RoleID = input.RoleID

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Update Successfully",
		"user":    user,
	})

}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	userID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID!"})
		return
	}

	if err := database.DB.Delete(&userModel.User{}, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}

func GetDeletedUsers(c *gin.Context) {
	var users []userModel.User
	database.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&users)
	c.JSON(http.StatusOK, users)
}

func RestoreUser(c *gin.Context) {
	userID := c.Param("id")
	database.DB.Unscoped().Model(&userModel.User{}).Where("id = ?", userID).Update("deleted_at", nil)
	c.JSON(http.StatusOK, gin.H{"message": "User restored successfully"})
}
