package controllers

import (
	"net/http"

	user_controller "delivery_app_api.mmedic.com/m/v2/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, uc *user_controller.UserController) {
	router.GET("/", HelloWorld)
	router.POST("/register", uc.Register)
	router.POST("/login", uc.Login)
}

func HelloWorld(c *gin.Context) {
	c.Status(http.StatusOK)
}
