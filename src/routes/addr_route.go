package routes

import (
	"delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"github.com/gin-gonic/gin"
)

func SetupAddrRoutes(router *gin.Engine, ac *controllers.AddrController) {
	router.GET("/address", authentication_utils.Authenticate("CUSTOMER"), ac.GetCustomerAddr)
}
