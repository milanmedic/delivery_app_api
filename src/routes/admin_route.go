package routes

import (
	"delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(router *gin.Engine, ac *controllers.AdminController) {
	router.GET("/admin/profile/:id", authentication_utils.Authenticate("ADMIN"), ac.GetAdminInfo)
	router.POST("/admin/login", ac.AdminLogin)
	router.POST("/admin/deliverer/create", authentication_utils.Authenticate("ADMIN"), ac.AddDeliverer)
	router.PATCH("/admin/customer/verify", authentication_utils.Authenticate("ADMIN"), ac.VerifyCustomer)
	router.PATCH("/admin/deliverer/verify", authentication_utils.Authenticate("ADMIN"), ac.VerifyDeliverer)
	router.POST("/admin/profile/update/:id", authentication_utils.Authenticate("ADMIN"), ac.UpdateAdmin)
}
