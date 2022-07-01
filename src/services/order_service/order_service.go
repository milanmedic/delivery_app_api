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

func (os *OrderService) GetOrdersByUserId(id string) ([]models.Order, error) {
	return os.repository.GetOrdersByUserId(id)
}

func (os *OrderService) CancelOrder(id string) error {
	return os.repository.UpdateProperty("status", "CANCELLED", id)
}

func (os *OrderService) GetOrderStatus(id string) (string, error) {
	return os.repository.GetOrderStatus(id)
}

func (os *OrderService) GetOrderBasketID(id string) (string, error) {
	return os.repository.GetOrderBasketID(id)
}

func (os *OrderService) GetAllOrders(deliveryStatus string, accepted ...string) ([]models.Order, error) {
	return os.repository.GetAllOrders(deliveryStatus, accepted...)
}

func (os *OrderService) AcceptOrder(orderID, delivererID string) error {
	err := os.repository.UpdateProperty("deliverer", delivererID, orderID)
	if err != nil {
		return err
	}
	err = os.repository.UpdateProperty("accepted", true, orderID)
	if err != nil {
		return err
	}

	return nil
}
