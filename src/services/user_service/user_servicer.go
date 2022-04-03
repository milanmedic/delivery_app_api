package services

import (
	dto "delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type UserServicer interface {
	GetUser(attr string, value interface{}) (*models.User, error)
	Exists(email string) (bool, error)
	ValidateUserRegistrationInput(udto dto.UserInputDto) error
	CreateUser(ud dto.UserInputDto) error
}
