package order_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/order_repository"
	"delivery_app_api.mmedic.com/m/v2/src/services/basket_service"
	"github.com/google/uuid"
)

type OrderService struct {
	repository    order_repository.OrderRepositer
	basketService *basket_service.BasketService
}

func CreateOrderService(or order_repository.OrderRepositer, bs *basket_service.BasketService) *OrderService {
	return &OrderService{repository: or, basketService: bs}
}

func (os *OrderService) CreateOrder(odto dto.OrderInputDto) error {
	//1. Use BasketService add Basket
	// 1.a. Using the Basket service, link articles to basket in the article_basket table

	//2. Afterwards call the OrderService to create a new order
	odto.Basket.Id = uuid.NewString()
	err := os.basketService.AddBasket(odto.Basket)
	if err != nil {
		return err
	}

	//TODO: If order creation fails, delete the basket from the database
	err = os.repository.CreateOrder(odto)
	if err != nil {
		return err
	}

	return nil
}
