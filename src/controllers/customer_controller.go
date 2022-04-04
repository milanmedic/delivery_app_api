package controllers

import (
	"fmt"
	"net/http"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	addr_service "delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	customer_service "delivery_app_api.mmedic.com/m/v2/src/services/customer_service"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService customer_service.CustomerServicer
	addrService     addr_service.AddrServicer
}

func CreateCustomerController(customerService customer_service.CustomerServicer, addrService addr_service.AddrServicer) *CustomerController {
	return &CustomerController{customerService: customerService, addrService: addrService}
}

func (uc *CustomerController) Register(c *gin.Context) {
	var customerDto dto.CustomerInputDto
	err := c.BindJSON(&customerDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = uc.customerService.ValidateCustomerRegistrationInput(customerDto)
	if err != nil {
		c.Error(fmt.Errorf("Error while validating input. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var addrId int
	addr, err := uc.addrService.GetAddr(*customerDto.Address)
	if err != nil {
		c.Error(fmt.Errorf("Error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if addr == nil {
		addrId, err = uc.addrService.CreateAddress(*customerDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		customerDto.Address.Id = addrId
	} else {
		customerDto.Address.Id = addr.Id
	}

	exists, err := uc.customerService.Exists(customerDto.Email)
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
		err = uc.customerService.CreateCustomer(customerDto)
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

func (uc *CustomerController) Login(c *gin.Context) {
	c.String(200, "Login")
}
