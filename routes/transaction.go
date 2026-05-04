package routes

import (
	"github.com/MouslyCode/bang-cukur/controllers"
	"github.com/MouslyCode/bang-cukur/middleware"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.Engine) {
	transaction := r.Group("/transaction", middleware.AuthMiddleware())
	transaction.GET("", middleware.AuthMiddleware(), controllers.GetTransactions)
	transaction.GET("/:id", middleware.AuthMiddleware(), controllers.GetTransactionById)
	transaction.POST("", middleware.AuthMiddleware(), controllers.CreateTransaction)
	transaction.DELETE("/:id", middleware.AuthMiddleware(), controllers.DeleteTransaction)
	transaction.GET("/deleted", middleware.AuthMiddleware(), controllers.GetDeletedTransactions)
	transaction.PUT("/restore/:id", middleware.AuthMiddleware(), controllers.RestoreTransaction)
}
