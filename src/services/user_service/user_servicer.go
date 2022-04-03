package services

import (
	dto "delivery_app_api.mmedic.com/m/v2/src/dto"
)

type UserServicer interface {
	CreateUser(ud dto.UserInputDto) error
}
