package addr_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type AddrServicer interface {
	GetById(id string) (*models.Address, error)
	CreateAddress(a dto.AddressInputDto) (int, error)
	GetAddr(a dto.AddressInputDto) (*models.Address, error)
	GetUserAddress(userId string) (*dto.AddressOutputDto, error)
}
