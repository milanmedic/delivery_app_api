package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}
	article, err := ac.articleService.GetBy("name", articleDto.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error while creating article")
		return
	}

	if article != nil {
		c.JSON(http.StatusBadRequest, "Article already exists.")
		return
	}

	err = ac.articleService.CreateArticle(articleDto)
	if err != nil {
		c.Error(fmt.Errorf("error while creating the article. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "Created")
}

func (ac *ArticleController) GetArticle(c *gin.Context) {
	id := c.Param("id")

	article, err := ac.articleService.GetBy("id", id)
	if err != nil {
		c.Error(fmt.Errorf("error while searching for the article. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if article == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, article)
}

func (ac *ArticleController) GetAllArticles(c *gin.Context) {
	articles, err := ac.articleService.GetAll()
	if err != nil {
		c.Error(fmt.Errorf("error while getting the articles. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if len(articles) <= 0 {
		c.Status(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, articles)
}

func (ac *ArticleController) DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error happened while searhing for article")
		return
	}

	article, err := ac.articleService.GetBy("id", id)
	if err != nil {
		c.Error(fmt.Errorf("error while deleting the article \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if article == nil {
		c.JSON(http.StatusInternalServerError, "Article doesn't exist")
		return
	}

	err = ac.articleService.DeleteArticle(id)
	if err != nil {
		c.Error(fmt.Errorf("error while deleting the article \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Deleted.")
}
