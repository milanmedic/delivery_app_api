package routes

import (
	"delivery_app_api.mmedic.com/m/v2/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, oc *controllers.OrderController) {
	router.POST("/order", oc.CreateOrder)
}
