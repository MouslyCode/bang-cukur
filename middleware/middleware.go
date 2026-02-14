package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/MouslyCode/bang-cukur/common/helper"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")

		claims, err := helper.VerifyJwt(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("user_id", claims.UserId)
		c.Set("role_id", claims.RoleId)

		c.Next()
	}
}

func RoleOnly(roles ...uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role_id")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Role Not Found"})
			return
		}

		roleID, ok := roleVal.(uint)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid role type"})
		}

		if slices.Contains(roles, roleID) {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
	}
}
