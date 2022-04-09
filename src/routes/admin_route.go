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

	router.POST("/admin/deliverer", authentication_utils.Authenticate("ADMIN"), ac.AddDeliverer)
	router.PATCH("/admin/customer/verification", authentication_utils.Authenticate("ADMIN"), ac.VerifyCustomer)
	router.PATCH("/admin/deliverer/verification", authentication_utils.Authenticate("ADMIN"), ac.VerifyDeliverer)
}
