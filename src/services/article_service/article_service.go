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

	return as.repository.AddOne(*article)
}

func (as *ArticleService) GetBy(attr string, value interface{}) (*models.Article, error) {
	return as.repository.GetBy(attr, value)
}

func (as *ArticleService) GetAll() ([]*models.Article, error) {
	return as.repository.GetAll()
}