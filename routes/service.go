package routes

import (
	"github.com/MouslyCode/bang-cukur/common/constant"
	"github.com/MouslyCode/bang-cukur/controllers"
	"github.com/MouslyCode/bang-cukur/middleware"
	"github.com/gin-gonic/gin"
)

func ServiceRoutes(r *gin.Engine) {
	service := r.Group("/service", middleware.AuthMiddleware())

	service.GET("", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetService)
	service.POST("", middleware.RoleOnly(constant.RoleOwnerID), controllers.CreateService)
	service.PUT("/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.UpdateService)
	service.DELETE("/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.DeleteService)
	service.GET("/deleted", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetDeletedServices)
	service.PATCH("/restore/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.RestoreService)
}
