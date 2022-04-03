package services

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	user_repo "delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/user_repository"
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
	//TODO: HASH PASSWORD BEFORE SAVING
	user.SetPassword(ud.Password)
	user.SetDateOfBirth(ud.DateOfBirth)
	user.SetAddress(addr)
	user.SetRole("CUSTOMER")
	user.SetVerificationStatus("UNVERIFIED")

	return us.repository.CreateUser(user)
}
