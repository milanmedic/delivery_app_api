package controllers

import (
	"net/http"

	customer_controller "delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.Engine, uc *customer_controller.CustomerController) {
	router.GET("/", HelloWorld)
	router.POST("/register", uc.Register)
	router.POST("/login", uc.Login)
	router.GET("/protected", authentication_utils.Authenticate(), HelloWorld)
	router.GET("/refresh", authentication_utils.Authenticate(), authentication_utils.RefreshToken)
}

func HelloWorld(c *gin.Context) {
	c.Status(http.StatusOK)
}
