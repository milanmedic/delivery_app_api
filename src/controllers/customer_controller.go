package controllers

import (
	"fmt"
	"net/http"
	"time"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	addr_service "delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	customer_service "delivery_app_api.mmedic.com/m/v2/src/services/customer_service"
	"delivery_app_api.mmedic.com/m/v2/src/utils/env_utils"
	"delivery_app_api.mmedic.com/m/v2/src/utils/jwt_utils"
	"delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"delivery_app_api.mmedic.com/m/v2/src/utils/validations"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

	user, err := uc.customerService.GetCustomer("email", credentials.Email)
	if err != nil {
		c.Error(fmt.Errorf("Error while creating a customer. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		c.String(http.StatusNotFound, "User with the provided email was not found.")
		return
	}

	if !security.CheckPasswordHash(credentials.Password, user.Password) {
		c.String(http.StatusUnauthorized, "Wrong password.")
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &jwt_utils.Claims{
		Email: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	jwtKey := []byte(env_utils.GetEnvVar("SECRET"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}
