package routes

import (
	"github.com/MouslyCode/bang-cukur/common/constant"
	"github.com/MouslyCode/bang-cukur/controllers"
	"github.com/MouslyCode/bang-cukur/middleware"
	"github.com/gin-gonic/gin"
)

func ItemRoutes(r *gin.Engine) {
	item := r.Group("/item", middleware.AuthMiddleware())

	item.GET("", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetItems)
	item.GET("/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetItemByID)
	item.GET("/:type", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetItemByType)
	item.POST("", middleware.RoleOnly(constant.RoleOwnerID), controllers.CreateItem)
	item.PUT("/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.UpdateItem)
	item.DELETE("/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.DeleteItem)
	item.GET("/deleted", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetDeletedItems)
	item.PUT("/restore/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.RestoreItem)
}
