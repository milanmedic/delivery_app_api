package controllers

import (
	"net/http"

	customer_controller "delivery_app_api.mmedic.com/m/v2/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.Engine, uc *customer_controller.CustomerController) {
	router.GET("/", HelloWorld)
	router.POST("/register", uc.Register)
	router.POST("/login", uc.Login)
}

func HelloWorld(c *gin.Context) {
	c.Status(http.StatusOK)
}
