package routes

import (
	"delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(router *gin.Engine, ac *controllers.AdminController) {
	router.PATCH("/verify", authentication_utils.Authenticate("ADMIN"), ac.VerifyCustomer)
}
