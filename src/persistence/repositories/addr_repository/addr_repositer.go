package addrrepository

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type AddrRepositer interface {
	GetById(id string) (*models.Address, error)
	CreateAddr(a dto.AddressInputDto) (int, error)
}
