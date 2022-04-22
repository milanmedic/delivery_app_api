package order_sql_db

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type OrderDber interface {
	CreateOrder(odto dto.OrderInputDto) error
	GetOrdersByUserId(id string) ([]models.Order, error)
}
