package routes

import (
	"github.com/MouslyCode/bang-cukur/common/constant"
	"github.com/MouslyCode/bang-cukur/controllers"
	"github.com/MouslyCode/bang-cukur/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	user := r.Group("/user", middleware.AuthMiddleware())

	user.GET("", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetUser)
	user.POST("", middleware.RoleOnly(constant.RoleOwnerID), controllers.CreateUser)
	user.PUT("/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.UpdateUser)
	user.DELETE("/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.DeleteUser)
	user.GET("/deleted", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetDeletedUsers)
	user.PUT("/restore/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.RestoreUser)
}
