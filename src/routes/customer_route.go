package routes

import (
	controllers "delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.Engine, uc *controllers.CustomerController) {
	router.POST("/customer/login", uc.Login)
	router.POST("/customer/registration", uc.Register)

	router.GET("/customer/login/github", uc.SendLoginOAuthRequest)
	router.GET("/customer/login/github/callback", uc.OAuthLogin)
	router.GET("/customer/register/github", uc.SendRegistrationOAuthRequest)
	router.GET("/customer/register/github/callback", uc.OAuthRegistration)

	router.GET("/customer/:id", authentication_utils.Authenticate("CUSTOMER"), uc.GetCustomerInfo)
	router.PUT("/customer/:id", authentication_utils.Authenticate("CUSTOMER"), uc.UpdateCustomer)
}
