package deliverer_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/deliverer_repository"
	"delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"delivery_app_api.mmedic.com/m/v2/src/utils/validations"
	"github.com/google/uuid"
)

type DelivererService struct {
	delivererRepository deliverer_repository.DelivererRepositer
}

func CreateDelivererService(dr deliverer_repository.DelivererRepositer) *DelivererService {
	return &DelivererService{delivererRepository: dr}
}

func (ds *DelivererService) AddDeliverer(ddto dto.DelivererInputDto) error {
	var deliverer models.Deliverer
	var addr *models.Address = models.CreateAddress(ddto.Address.Id, ddto.Address.StreetNum, ddto.Address.City, ddto.Address.Street, ddto.Address.Postfix)

	deliverer.SetId(uuid.NewString())
	deliverer.SetUsername(ddto.Username)
	deliverer.SetName(ddto.Name)
	deliverer.SetSurname(ddto.Surname)
	deliverer.SetEmail(ddto.Email)
	hash, err := security.HashPassword(ddto.Password)
	if err != nil {
		return err
	}
	deliverer.SetPassword(hash)
	deliverer.SetDateOfBirth(ddto.DateOfBirth)
	deliverer.SetAddress(addr)
	deliverer.SetRole("DELIVERER")
	deliverer.SetDeliveryProgress(false)
	deliverer.SetVerificationStatus("UNVERIFIED")

	return ds.delivererRepository.AddDeliverer(deliverer)
}

func (ds *DelivererService) GetBy(attr string, value interface{}) (*models.Deliverer, error) {
	return ds.delivererRepository.GetBy(attr, value)
}

func (ds *DelivererService) Exists(email string) (bool, error) {
	user, err := ds.GetBy("email", email)
	if err != nil {
		return false, err
	}
	if user != nil {
		return true, nil
	}

	return false, nil
}
func (ds *DelivererService) ValidateDelivererRegistrationInput(udto dto.DelivererInputDto) error {
	err := validations.ValidateName(udto.Name)
	if err != nil {
		return err
	}
	err = validations.ValidateSurname(udto.Surname)
	if err != nil {
		return err
	}
	err = validations.ValidateUsername(udto.Username)
	if err != nil {
		return err
	}
	err = validations.ValidatePassword(udto.Password)
	if err != nil {
		return err
	}
	err = validations.ValidateEmail(udto.Email)
	if err != nil {
		return err
	}
	err = validations.ValidateCity(udto.Address.City)
	if err != nil {
		return err
	}
	err = validations.ValidatePostfix(udto.Address.Postfix)
	if err != nil {
		return err
	}
	err = validations.ValidateStreetNum(udto.Address.StreetNum)
	if err != nil {
		return err
	}
	err = validations.ValidateStreet(udto.Address.Street)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DelivererService) UpdateProperty(property string, value interface{}, id string) error {
	return ds.delivererRepository.UpdateProperty(property, value, id)
}

func (ds *DelivererService) GetDelivererInfo(id string) (*dto.DelivererOutputDto, error) {
	deliverer, err := ds.GetBy("id", id)
	if err != nil {
		return nil, err
	}

	if deliverer == nil {
		return nil, nil
	}

	var delivererOutputDto *dto.DelivererOutputDto = new(dto.DelivererOutputDto)
	var addressOutputDto *dto.AddressOutputDto = new(dto.AddressOutputDto)

	addressOutputDto.City = deliverer.Address.City
	addressOutputDto.Street = deliverer.Address.Street
	addressOutputDto.StreetNum = deliverer.Address.StreetNum
	addressOutputDto.Postfix = deliverer.Address.Postfix

	delivererOutputDto.Address = addressOutputDto
	delivererOutputDto.Name = deliverer.Name
	delivererOutputDto.Surname = deliverer.Surname
	delivererOutputDto.Email = deliverer.Email
	delivererOutputDto.Username = deliverer.Username
	delivererOutputDto.DateOfBirth = deliverer.DateOfBirth
	delivererOutputDto.DeliveryInProgess = deliverer.DeliveryInProgress

	return delivererOutputDto, nil
}
