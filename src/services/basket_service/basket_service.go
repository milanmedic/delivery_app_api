package basket_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/basket_repository"
)

type BasketService struct {
	repository *basket_repository.BasketRepository
}

func CreateBasketService(br *basket_repository.BasketRepository) *BasketService {
	return &BasketService{repository: br}
}

func (bs *BasketService) AddBasket(bdto dto.BasketInputDto) error {
	return bs.repository.AddBasket(bdto)
}
