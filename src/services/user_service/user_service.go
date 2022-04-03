package services

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	user_repo "delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/user_repository"
	"delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"delivery_app_api.mmedic.com/m/v2/src/utils/validations"
	"github.com/google/uuid"
)

type UserService struct {
	repository user_repo.UserRepositer
}

func CreateUserService(repo user_repo.UserRepositer) *UserService {
	return &UserService{repository: repo}
}

func (us *UserService) CreateUser(ud dto.UserInputDto) error {
	var user models.User
	var addr *models.Address = models.CreateAddress(ud.Address.Id, ud.Address.StreetNum, ud.Address.City, ud.Address.Street, ud.Address.Postfix)

	user.SetId(uuid.NewString())
	user.SetUsername(ud.Username)
	user.SetName(ud.Name)
	user.SetSurname(ud.Surname)
	user.SetEmail(ud.Email)
	hash, err := security.HashPassword(ud.Password)
	if err != nil {
		return err
	}
	user.SetPassword(hash)
	user.SetDateOfBirth(ud.DateOfBirth)
	user.SetAddress(addr)
	user.SetRole("CUSTOMER")
	user.SetVerificationStatus("UNVERIFIED")

	return us.repository.CreateUser(user)
}

func (us *UserService) ValidateUserRegistrationInput(udto dto.UserInputDto) error {
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

func (us *UserService) GetUser(attr string, value interface{}) (*models.User, error) {
	return us.repository.GetUser(attr, value)
}

func (us *UserService) Exists(email string) (bool, error) {
	user, err := us.GetUser("email", email)
	if err != nil {
		return false, err
	}
	if user != nil {
		return true, nil
	}

	return false, nil
}
