package routes

import (
	"delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"github.com/gin-gonic/gin"
)

func SetupDelivererRoutes(router *gin.Engine, dc *controllers.DelivererController) {
	router.POST("/deliverer/login", dc.DelivererLogin)
	router.GET("/deliverer/:id", authentication_utils.Authenticate("DELIVERER"), dc.GetDelivererInfo)

	router.PUT("/deliverer/:id", authentication_utils.Authenticate("DELIVERER"), dc.UpdateDeliverer)
}
