package article_repository

import (
	"delivery_app_api.mmedic.com/m/v2/src/models"
	article_sql_db "delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/article_db_impls"
)

type ArticleRepository struct {
	db article_sql_db.ArticleDber
}

func CreateArticleRepository(db article_sql_db.ArticleDber) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (ar *ArticleRepository) GetBy(attr string, value interface{}) (*models.Article, error) {
	return ar.db.GetBy(attr, value)
}

func (ar *ArticleRepository) AddOne(a models.Article) error {
	return ar.db.AddOne(a)
}

func (ar *ArticleRepository) GetAll() ([]*models.Article, error) {
	return ar.db.GetAll()
}

func (ac *ArticleRepository) UpdateProperty(property string, value interface{}, id int) error {
	return ac.db.UpdateProperty(property, value, id)
}

func (ac *ArticleRepository) DecrementQuantity(value int, id int) error {
	return ac.db.DecrementQuantity(value, id)
}

func (ac *ArticleRepository) DeleteArticle(id int) error {
	return ac.db.DeleteArticle(id)
}
