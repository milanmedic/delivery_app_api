package routes

import (
	"delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(router *gin.Engine, ac *controllers.AdminController) {
	router.POST("/admin/login", ac.AdminLogin)

	router.GET("/admin/:id", authentication_utils.Authenticate("ADMIN"), ac.GetAdminInfo)
	router.PUT("/admin/:id", authentication_utils.Authenticate("ADMIN"), ac.UpdateAdmin)

	router.GET("/admin/deliverers", authentication_utils.Authenticate("ADMIN"), ac.GetAllDeliverers)

	router.PATCH("/admin/deliverer/verification", authentication_utils.Authenticate("ADMIN"), ac.VerifyDeliverer)
}
