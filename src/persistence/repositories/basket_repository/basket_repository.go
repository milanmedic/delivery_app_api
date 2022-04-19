package basket_repository

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/basket_sql_db"
)

type BasketRepository struct {
	db basket_sql_db.BasketDber
}

func CreateBasketRepository(db basket_sql_db.BasketDber) *BasketRepository {
	return &BasketRepository{db: db}
}

func (br *BasketRepository) AddBasket(bdto dto.BasketInputDto) error {
	return br.db.AddBasket(bdto)
}
