package order_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type OrderServicer interface {
	CreateOrder(odto dto.OrderInputDto) error
	GetOrdersByUsername(username string) ([]models.Order, error)
}
