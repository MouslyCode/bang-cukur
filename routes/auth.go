package routes

import (
	"github.com/MouslyCode/bang-cukur/common/constant"
	"github.com/MouslyCode/bang-cukur/controllers"
	"github.com/MouslyCode/bang-cukur/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	user := r.Group("/", middleware.AuthMiddleware())

	user.POST("/create", middleware.RoleOnly(constant.RoleOwnerID), controllers.UserCreate)
}
