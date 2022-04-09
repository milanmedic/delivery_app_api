package controllers

import (
	"fmt"
	"net/http"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/services/article_service"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	articleService article_service.ArticleServicer
}

func CreateArticleController(as article_service.ArticleServicer) *ArticleController {
	return &ArticleController{articleService: as}
}

func (ac *ArticleController) CreateArticle(c *gin.Context) {
	var articleDto dto.ArticleInputDto
	err := c.BindJSON(&articleDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// TODO: Data Validation
	err = ac.articleService.CreateArticle(articleDto)
	if err != nil {
		c.Error(fmt.Errorf("Error while creating the article. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
	return
}

func (ac *ArticleController) GetArticle(c *gin.Context) {
	id := c.Param("id")

	article, err := ac.articleService.GetBy("id", id)
	if err != nil {
		c.Error(fmt.Errorf("Error while searching for the article. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if article == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, article)
	return
}

func (ac *ArticleController) GetAllArticles(c *gin.Context) {

	articles, err := ac.articleService.GetAll()
	if err != nil {
		c.Error(fmt.Errorf("Error while getting the articles. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if len(articles) <= 0 {
		c.Status(http.StatusNotFound)
	}

	c.JSON(200, articles)
	return
}
