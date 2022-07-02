package article_repository

import "delivery_app_api.mmedic.com/m/v2/src/models"

type ArticleRepositer interface {
	GetBy(attr string, value interface{}) (*models.Article, error)
	AddOne(a models.Article) error
	GetAll() ([]*models.Article, error)
	UpdateProperty(property string, value interface{}, id int) error
	DecrementQuantity(value int, id int) error
	DeleteArticle(id int) error
}
