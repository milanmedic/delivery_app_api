package order_sql_db

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type OrderDber interface {
	CreateOrder(odto dto.OrderInputDto) error
	GetOrdersByUserId(id string) ([]models.Order, error)
	DeleteOrder(attr string, value interface{}) error
	GetOrderBasketID(id string) (string, error)
	GetOrderStatus(id string) (string, error)
	UpdateProperty(property string, value interface{}, id string) error
	GetAllOrders(deliveryStatus string, accepted ...string) ([]models.Order, error)
}
