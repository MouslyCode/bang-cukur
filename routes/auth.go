package routes

import (
	"github.com/MouslyCode/bang-cukur/controllers"
	"github.com/MouslyCode/bang-cukur/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	user := r.Group("/")
	{
		user.POST("/login", controllers.Login)
		user.POST("/create", middleware.OwnerMiddleware(), controllers.UserCreate)
	}
}
