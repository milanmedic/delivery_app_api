package controllers

import (
	"fmt"
	"net/http"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	addr_service "delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	user_service "delivery_app_api.mmedic.com/m/v2/src/services/user_service"
	"github.com/gin-gonic/gin"
)

//TODO: Change all User related naming from USER -> CUSTOMER
type UserController struct {
	userService user_service.UserServicer
	addrService addr_service.AddrServicer
}

func CreateUserController(userService user_service.UserServicer, addrService addr_service.AddrServicer) *UserController {
	return &UserController{userService: userService, addrService: addrService}
}

func (uc *UserController) Register(c *gin.Context) {
	var userDto dto.UserInputDto
	err := c.BindJSON(&userDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = uc.userService.ValidateUserRegistrationInput(userDto)
	if err != nil {
		c.Error(fmt.Errorf("Error while validating input. \nReason: %s", err.Error()))
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var addrId int

	if userDto.Address.Id == 0 {
		addrId, err = uc.addrService.CreateAddress(*userDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	userDto.Address.Id = addrId

	exists, err := uc.userService.Exists(userDto.Email)
	if err != nil {
		c.Error(fmt.Errorf("Error while checking for a customer. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		c.String(http.StatusBadRequest, "User with the provided email already exists")
		return
	}

	if !exists {
		err = uc.userService.CreateUser(userDto)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating a customer. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.Status(204)
	return
}

func (uc *UserController) Login(c *gin.Context) {
	c.String(200, "Login")
}
