package middleware

import (
	"net/http"
	"strings"

	"github.com/MouslyCode/bang-cukur/common/constant"
	"github.com/MouslyCode/bang-cukur/common/helper"
	"github.com/gin-gonic/gin"
)

func OwnerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		c.Set("token", token)

		userId, err := helper.VerifyJwt(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("user_id", userId)
		roleId := c.GetInt("role_id")
		if roleId != constant.RoleOwnerID {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Prohibited to access"})
			c.Abort()
			return
		}
		c.Next()

	}
}

func CashierMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		c.Set("token", token)
		c.Next()

	}
}
