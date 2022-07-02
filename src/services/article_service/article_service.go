package article_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/article_repository"
)

type ArticleService struct {
	repository article_repository.ArticleRepositer
}

func CreateArticleService(ar article_repository.ArticleRepositer) *ArticleService {
	return &ArticleService{repository: ar}
}

func (as *ArticleService) CreateArticle(adto dto.ArticleInputDto) error {
	var article *models.Article = new(models.Article)

	article.Name = adto.Name
	article.Description = adto.Description
	article.Price = adto.Price
	article.Quantity = adto.Quantity

	return as.repository.AddOne(*article)
}

func (as *ArticleService) GetBy(attr string, value interface{}) (*models.Article, error) {
	return as.repository.GetBy(attr, value)
}

func (as *ArticleService) GetAll() ([]*models.Article, error) {
	return as.repository.GetAll()
}

func (as *ArticleService) UpdateProperty(property string, value interface{}, id int) error {
	return as.repository.UpdateProperty(property, value, id)
}

func (as *ArticleService) DecrementQuantity(value int, id int) error {
	return as.repository.DecrementQuantity(value, id)
}

func (as *ArticleService) DeleteArticle(id int) error {
	return as.repository.DeleteArticle(id)
}
