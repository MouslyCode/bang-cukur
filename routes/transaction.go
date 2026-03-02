package routes

import (
	"github.com/MouslyCode/bang-cukur/controllers"
	"github.com/MouslyCode/bang-cukur/middleware"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.Engine) {
	transaction := r.Group("/transaction", middleware.AuthMiddleware())

	transaction.POST("", controllers.CreateTransaction)
}
