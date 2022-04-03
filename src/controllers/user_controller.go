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
	}

	var addrId int
	// IF address_id empty
	// it means the address needs to be created
	// create a new address, and retireve it's id
	if userDto.Address.Id == 0 {
		addrId, err = uc.addrService.CreateAddress(*userDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating an address. \nReason: %s", err.Error()))
			c.Status(http.StatusInternalServerError)
		}
	}

	userDto.Address.Id = addrId

	//TODO: CHECK IF USER EXISTS BEFORE CREATING A NEW USER

	err = uc.userService.CreateUser(userDto)
	if err != nil {
		c.Error(fmt.Errorf("Error while creating a customer. \nReason: %s", err.Error()))
		c.Status(http.StatusInternalServerError)
	}

	c.Status(204)
}

func (uc *UserController) Login(c *gin.Context) {
	c.String(200, "Login")
}
