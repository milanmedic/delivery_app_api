package routes

import (
	controllers "delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.Engine, uc *controllers.CustomerController) {
	router.POST("/register", uc.Register)
	router.POST("/login", uc.Login)
	router.GET("/refresh", authentication_utils.Authenticate("*"), authentication_utils.RefreshToken)
	router.GET("login/github", uc.SendLoginOAuthRequest)
	router.GET("login/github/callback", uc.OAuthLogin)
	router.GET("register/github", uc.SendRegistrationOAuthRequest)
	router.GET("register/github/callback", uc.OAuthRegistration)
}
