package admin_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/admin_repository"
	"delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"delivery_app_api.mmedic.com/m/v2/src/utils/validations"
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

func (as *AdminService) UpdateAdmin(id string, adto *dto.AdminInputDto) (bool, error) {
	admin, err := as.GetBy("id", id)
	if err != nil {
		return false, err
	}
	if admin == nil {
		return false, nil
	}

	admin.SetName(adto.Name)
	admin.SetSurname(adto.Surname)
	admin.SetUsername(adto.Username)
	admin.SetEmail(adto.Email)
	hash, err := security.HashPassword(adto.Password)
	if err != nil {
		return false, err
	}
	admin.SetPassword(hash)
	admin.SetDateOfBirth(adto.DateOfBirth)
	admin.SetAddress((*models.Address)(adto.Address))

	return as.repository.Update(admin)
}

func (as *AdminService) ValidateAdminDataInput(adto dto.AdminInputDto) error {
	err := validations.ValidateName(adto.Name)
	if err != nil {
		return err
	}
	err = validations.ValidateSurname(adto.Surname)
	if err != nil {
		return err
	}
	err = validations.ValidateUsername(adto.Username)
	if err != nil {
		return err
	}
	err = validations.ValidatePassword(adto.Password)
	if err != nil {
		return err
	}
	err = validations.ValidateEmail(adto.Email)
	if err != nil {
		return err
	}
	err = validations.ValidateCity(adto.Address.City)
	if err != nil {
		return err
	}
	err = validations.ValidatePostfix(adto.Address.Postfix)
	if err != nil {
		return err
	}
	err = validations.ValidateStreetNum(adto.Address.StreetNum)
	if err != nil {
		return err
	}
	err = validations.ValidateStreet(adto.Address.Street)
	if err != nil {
		return err
	}

	return nil
}
