package addr_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	addr_repository "delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/addr_repository"
	"delivery_app_api.mmedic.com/m/v2/src/utils/validations"
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

func (as *AddrService) GetAddr(a dto.AddressInputDto) (*models.Address, error) {
	return as.repository.GetAddr(a)
}

func (as *AddrService) GetById(id string) (*models.Address, error) {
	return as.repository.GetById(id)
}

func (as *AddrService) GetUserAddress(userId string) (*dto.AddressOutputDto, error) {
	return as.repository.GetUserAddress(userId)
}

func (as *AddrService) ValidateAddress(city, postfix, street string, streetNum int) error {
	err := validations.ValidateCity(city)
	if err != nil {
		return err
	}
	err = validations.ValidatePostfix(postfix)
	if err != nil {
		return err
	}
	err = validations.ValidateStreetNum(streetNum)
	if err != nil {
		return err
	}
	err = validations.ValidateStreet(street)
	if err != nil {
		return err
	}

	return nil
}
