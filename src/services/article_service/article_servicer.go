package article_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type ArticleServicer interface {
	CreateArticle(adto dto.ArticleInputDto) error
	GetBy(attr string, value interface{}) (*models.Article, error)
	GetAll() ([]*models.Article, error)
	UpdateProperty(property string, value interface{}, id int) error
	DecrementQuantity(value int, id int) error
	DeleteArticle(id int) error
}
