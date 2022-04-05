package customer_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	customer_repo "delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/customer_repository"
	"delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"delivery_app_api.mmedic.com/m/v2/src/utils/validations"
	"github.com/google/uuid"
)

type CustomerService struct {
	repository customer_repo.CustomerRepositer
}

func CreateCustomerService(repo customer_repo.CustomerRepositer) *CustomerService {
	return &CustomerService{repository: repo}
}

func (cs *CustomerService) CreateCustomer(ud dto.CustomerInputDto) error {
	var customer models.Customer
	var addr *models.Address = models.CreateAddress(ud.Address.Id, ud.Address.StreetNum, ud.Address.City, ud.Address.Street, ud.Address.Postfix)

	customer.SetId(uuid.NewString())
	customer.SetUsername(ud.Username)
	customer.SetName(ud.Name)
	customer.SetSurname(ud.Surname)
	customer.SetEmail(ud.Email)
	hash, err := security.HashPassword(ud.Password)
	if err != nil {
		return err
	}
	customer.SetPassword(hash)
	customer.SetDateOfBirth(ud.DateOfBirth)
	customer.SetAddress(addr)
	customer.SetRole("CUSTOMER")
	customer.SetVerificationStatus("UNVERIFIED")

	return cs.repository.CreateCustomer(customer)
}

func (cs *CustomerService) ValidateCustomerRegistrationInput(udto dto.CustomerInputDto) error {
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

func (cs *CustomerService) GetCustomer(attr string, value interface{}) (*models.Customer, error) {
	return cs.repository.GetCustomer(attr, value)
}

func (cs *CustomerService) Exists(email string) (bool, error) {
	user, err := cs.GetCustomer("email", email)
	if err != nil {
		return false, err
	}
	if user != nil {
		return true, nil
	}

	return false, nil
}

func (cs *CustomerService) UpdateProperty(property string, value interface{}, id string) error {
	return cs.repository.UpdateProperty(property, value, id)
}
