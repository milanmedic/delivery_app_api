package controllers

import (
	"fmt"
	"net/http"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	addr_service "delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	customer_service "delivery_app_api.mmedic.com/m/v2/src/services/customer_service"
	"delivery_app_api.mmedic.com/m/v2/src/utils/jwt_utils"
	"delivery_app_api.mmedic.com/m/v2/src/utils/oauth_utils"
	"delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"delivery_app_api.mmedic.com/m/v2/src/utils/validations"
	"github.com/gin-gonic/gin"
)

//TODO: Refactor
//TODO: Change all errors to custom error
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
			return
		}
	}

	c.Status(204)
	return
}

func (uc *CustomerController) Login(c *gin.Context) {
	var credentials jwt_utils.Credentials
	err := c.BindJSON(&credentials)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = validations.ValidateEmail(credentials.Email)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	customer, err := uc.customerService.GetBy("email", credentials.Email)
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if customer == nil {
		c.String(http.StatusNotFound, "User with the provided email was not found.")
		return
	}

	if !security.CheckPasswordHash(credentials.Password, customer.Password) {
		c.String(http.StatusUnauthorized, "Wrong password.")
		return
	}

	claims := jwt_utils.CreateClaims()
	claims.Email = credentials.Email
	claims.Role = customer.Role

	tokenString, err := jwt_utils.CreateToken(claims)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}

func (uc *CustomerController) SendLoginOAuthRequest(c *gin.Context) {
	reqURL := oauth_utils.GetLoginOAuthURL()
	c.Redirect(http.StatusFound, reqURL)
	return
}

func (uc *CustomerController) OAuthLogin(c *gin.Context) {
	code := c.Query("code")

	customerData, err := oauth_utils.GetCustomerGithubInformation(code, "LOGIN")
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	customer, err := uc.customerService.GetBy("email", customerData.Email)
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if customer == nil {
		c.String(http.StatusNotFound, "User with the provided email was not found.")
		return
	}

	claims := jwt_utils.CreateClaims()
	claims.Email = customer.Email
	claims.Role = customer.Role

	tokenString, err := jwt_utils.CreateToken(claims)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}

func (uc *CustomerController) SendRegistrationOAuthRequest(c *gin.Context) {
	reqURL := oauth_utils.GetRegistrationOAuthURL()
	c.Redirect(http.StatusFound, reqURL)
	return
}

func (uc *CustomerController) OAuthRegistration(c *gin.Context) {

	code := c.Query("code")

	customerData, err := oauth_utils.GetCustomerGithubInformation(code, "REGISTRATION")
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	exists, err := uc.customerService.Exists(customerData.Email)
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		c.String(http.StatusNotFound, "User with the provided email already exists.")
		return
	}

	addrId, err := uc.addrService.CreateAddress(*customerData.Address)
	if err != nil {
		c.Error(fmt.Errorf("Error while creating an address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	customerData.Address.Id = addrId

	if !exists {
		err = uc.customerService.CreateCustomer(*customerData)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating a customer. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Status(204)
	return
}

func (cc *CustomerController) GetCustomerInfo(c *gin.Context) {
	var id string = c.Param("id")

	adminOut, err := cc.customerService.GetCustomerInfo(id)
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if adminOut == nil {
		c.String(http.StatusNotFound, "Customer doesn't exist.")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": adminOut})
	return
}
