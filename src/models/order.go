package models

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
)

type Order struct {
	Id               string               `json:"id"`
	Comment          string               `json:"comment"`
	Accepted         bool                 `json:"accepted"`
	Status           string               `json:"status"` //PENDING, IN PROGRESS, COMPLETED, CANCELLED
	DelivererName    string               `json:"deliverer_name"`
	DelivererSurname string               `json:"deliverer_surname"`
	Address          dto.AddressOutputDto `json:"address"`
	Basket           dto.BasketOutputDto  `json:"basket"`
}
