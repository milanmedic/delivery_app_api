package order_repository

import "delivery_app_api.mmedic.com/m/v2/src/dto"

type OrderRepositer interface {
	CreateOrder(odto dto.OrderInputDto) error
}
