package routes

import (
	"github.com/MouslyCode/bang-cukur/common/constant"
	"github.com/MouslyCode/bang-cukur/controllers"
	"github.com/MouslyCode/bang-cukur/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	product := r.Group("/product", middleware.AuthMiddleware())

	product.GET("", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetItems)
	product.POST("", middleware.RoleOnly(constant.RoleOwnerID), controllers.CreateItem)
	product.PUT("/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.UpdateItem)
	product.DELETE("/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.DeleteItem)
	product.GET("/deleted", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetDeletedItems)
	product.PATCH("/restore/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.RestoreItem)
}
