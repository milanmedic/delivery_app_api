package routes

import (
	"delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"github.com/gin-gonic/gin"
)

func SetupDelivererRoutes(router *gin.Engine, dc *controllers.DelivererController) {
	router.POST("/deliverer/login", dc.DelivererLogin)
	router.POST("/deliverer/registration", dc.Register)
	router.GET("/deliverer", authentication_utils.Authenticate("DELIVERER"), dc.GetDelivererInfo)

	router.PUT("/deliverer/:id", authentication_utils.Authenticate("DELIVERER"), dc.UpdateDeliverer)
	router.PATCH("/deliverer/order/:id", authentication_utils.Authenticate("DELIVERER"), dc.AcceptOrder)

	router.PATCH("/deliverer", authentication_utils.Authenticate("DELIVERER"), dc.UpdateDelivererProperty)
	router.PATCH("/deliverer/password", authentication_utils.Authenticate("DELIVERER"), dc.UpdateDelivererPassword)
	router.PATCH("/deliverer/address", authentication_utils.Authenticate("DELIVERER"), dc.UpdateDelivererAddress)

	router.PATCH("/deliverer/order/completed/:id", authentication_utils.Authenticate("DELIVERER"), dc.CompleteOrder)

	router.GET("/delivery_time", dc.StreamDeliveryTime)
}
