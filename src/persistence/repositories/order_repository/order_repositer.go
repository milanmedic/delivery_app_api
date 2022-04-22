package order_repository

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type OrderRepositer interface {
	CreateOrder(odto dto.OrderInputDto) error
	GetOrdersBy(attr string, value interface{}) ([]models.Order, error)
}
