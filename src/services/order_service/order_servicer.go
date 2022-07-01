package order_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type OrderServicer interface {
	CreateOrder(odto dto.OrderInputDto) error
	GetOrdersByUserId(id string) ([]models.Order, error)
	GetAllOrders(deliveryStatus string, accepted ...string) ([]models.Order, error)
	CancelOrder(id string) error
	GetOrderBasketID(id string) (string, error)
	GetOrderStatus(id string) (string, error)
	AcceptOrder(orderID, delivererID string) error
}
