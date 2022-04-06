package admin_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/admin_repository"
)

type AdminService struct {
	repository admin_repository.AdminRepositer
}

func CreateAdminService(ar admin_repository.AdminRepositer) *AdminService {
	return &AdminService{repository: ar}
}

func (as *AdminService) GetBy(attr string, value interface{}) (*models.Admin, error) {
	return as.repository.GetBy(attr, value)
}

func (as *AdminService) GetAdminInfo(id string) (*dto.AdminOutputDto, error) {
	admin, err := as.GetBy("id", id)
	if err != nil {
		return nil, err
	}

	if admin == nil {
		return nil, nil
	}

	var adminOutputDto *dto.AdminOutputDto = new(dto.AdminOutputDto)
	var addressOutputDto *dto.AddressOutputDto = new(dto.AddressOutputDto)

	addressOutputDto.City = admin.Address.City
	addressOutputDto.Street = admin.Address.Street
	addressOutputDto.StreetNum = admin.Address.StreetNum
	addressOutputDto.Postfix = admin.Address.Postfix

	adminOutputDto.Address = addressOutputDto
	adminOutputDto.Name = admin.Name
	adminOutputDto.Surname = admin.Surname
	adminOutputDto.Email = admin.Email
	adminOutputDto.Username = admin.Username
	adminOutputDto.DateOfBirth = admin.DateOfBirth

	return adminOutputDto, nil
}
