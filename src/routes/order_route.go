package routes

import (
	"delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, oc *controllers.OrderController) {
	router.POST("/order", authentication_utils.Authenticate("CUSTOMER", "ADMIN"), oc.CreateOrder)
}
