package services

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	addr_repository "delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/addr_repository"
)

type AddrService struct {
	repository addr_repository.AddrRepositer
}

func CreateAddrService(ar addr_repository.AddrRepositer) *AddrService {
	return &AddrService{repository: ar}
}

func (as *AddrService) CreateAddress(a dto.AddressInputDto) (int, error) {
	return as.repository.CreateAddr(a)
}

func (as *AddrService) GetById(id string) (*models.Address, error) {
	return as.repository.GetById(id)
}
