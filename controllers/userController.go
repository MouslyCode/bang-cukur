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
	"gorm.io/gorm"
)

func UserCreate(c *gin.Context) {
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

	token, err := helper.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Login sucessfully",
		"token":   token,
	})

}
