package routes

import (
	"delivery_app_api.mmedic.com/m/v2/src/controllers"
	authentication_utils "delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"github.com/gin-gonic/gin"
)

func SetupArticleRoutes(router *gin.Engine, ac *controllers.ArticleController) {
	router.POST("/article", authentication_utils.Authenticate("ADMIN"), ac.CreateArticle)
	router.GET("/article/:id", ac.GetArticle)
	router.GET("/articles", ac.GetAllArticles)
}
