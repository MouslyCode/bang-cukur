package routes

import (
	"github.com/MouslyCode/bang-cukur/common/constant"
	"github.com/MouslyCode/bang-cukur/controllers"
	"github.com/MouslyCode/bang-cukur/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	product := r.Group("/product", middleware.AuthMiddleware())

	product.GET("", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetProducts)
	product.POST("", middleware.RoleOnly(constant.RoleOwnerID), controllers.CreateProduct)
	product.PUT("/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.UpdateProduct)
	product.DELETE("/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.DeleteProduct)
	product.GET("/deleted", middleware.RoleOnly(constant.RoleOwnerID), controllers.GetDeletedProducts)
	product.PATCH("/restore/:id", middleware.RoleOnly(constant.RoleOwnerID), controllers.RestoreProduct)
}
