package order_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
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
	odto.Basket.Id = uuid.NewString()
	err := os.basketService.AddBasket(odto.Basket)
	if err != nil {
		return err
	}

	err = os.repository.CreateOrder(odto)
	if err != nil {
		os.basketService.DeleteBasket(odto.Basket.Id)
		return err
	}

	return nil
}

func (os *OrderService) GetOrdersByUsername(username string) ([]models.Order, error) {
	return os.repository.GetOrdersBy("username", username)
}
