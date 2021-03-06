package routes

import (
	"delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, oc *controllers.OrderController) {
	router.POST("/order", authentication_utils.Authenticate("CUSTOMER"), oc.CreateOrder)
	router.GET("/orders", authentication_utils.Authenticate("CUSTOMER"), oc.GetOrders)
	router.GET("/orders/all", authentication_utils.Authenticate("DELIVERER", "ADMIN"), oc.GetAllOrders)
	router.PATCH("/order/:id", authentication_utils.Authenticate("CUSTOMER"), oc.CancelOrder)
	router.GET("/deliverer/orders", authentication_utils.Authenticate("DELIVERER"), oc.GetAllDelivererOrders)
}
