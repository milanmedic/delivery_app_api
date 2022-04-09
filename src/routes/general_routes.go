package routes

import (
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"github.com/gin-gonic/gin"
)

func SetupGeneralRoutes(router *gin.Engine) {
	router.GET("/refresh", authentication_utils.Authenticate("*"), authentication_utils.RefreshToken)
}
