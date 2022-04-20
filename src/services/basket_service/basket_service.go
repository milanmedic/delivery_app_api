package basket_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/basket_repository"
	"delivery_app_api.mmedic.com/m/v2/src/services/article_service"
)

type BasketService struct {
	repository     *basket_repository.BasketRepository
	articleService article_service.ArticleServicer
}

func CreateBasketService(br *basket_repository.BasketRepository, as article_service.ArticleServicer) *BasketService {
	return &BasketService{repository: br, articleService: as}
}

func (bs *BasketService) AddBasket(bdto dto.BasketInputDto) error {
	var err error
	for _, a := range bdto.Articles {
		err = bs.articleService.DecrementQuantity(a.Quantity, a.Id)
		if err != nil {
			return err
		}
	}
	return bs.repository.AddBasket(bdto)
}

func (bs *BasketService) DeleteBasket(bId string) error {
	return bs.repository.DeleteBasket(bId)
}
