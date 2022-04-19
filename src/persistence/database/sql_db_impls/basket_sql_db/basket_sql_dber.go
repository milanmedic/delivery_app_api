package basket_sql_db

import "delivery_app_api.mmedic.com/m/v2/src/dto"

type BasketDber interface {
	AddBasket(bdto dto.BasketInputDto) error
}
