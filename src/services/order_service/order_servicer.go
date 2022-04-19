package order_service

import "delivery_app_api.mmedic.com/m/v2/src/dto"

type OrderServicer interface {
	CreateOrder(odto dto.OrderInputDto) error
}
