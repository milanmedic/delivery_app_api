package order_sql_db

import "delivery_app_api.mmedic.com/m/v2/src/dto"

type OrderDber interface {
	CreateOrder(odto dto.OrderInputDto) error
}
