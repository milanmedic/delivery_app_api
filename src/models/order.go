package models

import "delivery_app_api.mmedic.com/m/v2/src/dto"

type Order struct {
	Id               string               `json:"id"`
	Comment          string               `json:"comment"`
	Status           bool                 `json:"status"`
	DelivererName    string               `json:"deliverer_name"`
	DelivererSurname string               `json:"deliverer_surname"`
	Address          dto.AddressOutputDto `json:"address"`
	Basket           dto.BasketOutputDto  `json:"basket"`
}
