package controllers

import (
	"fmt"
	"net/http"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
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

func (cc *CustomerController) Register(c *gin.Context) {
	var customerDto dto.CustomerInputDto
	err := c.BindJSON(&customerDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = cc.customerService.ValidateCustomerDataInput(customerDto)
	if err != nil {
		c.Error(fmt.Errorf("error while validating input. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var addrId int
	addr, err := cc.addrService.GetAddr(*customerDto.Address)
	if err != nil {
		c.Error(fmt.Errorf("error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if addr == nil {
		addrId, err = cc.addrService.CreateAddress(*customerDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		customerDto.Address.Id = addrId
	} else {
		customerDto.Address.Id = addr.Id
	}

	exists, err := cc.customerService.Exists(customerDto.Email)
	if err != nil {
		c.Error(fmt.Errorf("error while checking for a customer. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		c.String(http.StatusBadRequest, "User with the provided email already exists")
		return
	}

	if !exists {
		err = cc.customerService.CreateCustomer(customerDto)
		if err != nil {
			c.Error(fmt.Errorf("error while creating a customer. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusCreated, "Created.")
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
		c.Error(fmt.Errorf("error while retrieving the customer info. \nReason: %s", err.Error()))
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
	claims.UserId = customer.Id
	claims.Role = customer.Role

	tokenString, err := jwt_utils.CreateToken(claims)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (uc *CustomerController) SendLoginOAuthRequest(c *gin.Context) {
	reqURL := oauth_utils.GetLoginOAuthURL()
	c.Redirect(http.StatusFound, reqURL)
}

func (uc *CustomerController) OAuthLogin(c *gin.Context) {
	code := c.Query("code")

	customerData, err := oauth_utils.GetCustomerGithubInformation(code, "LOGIN")
	if err != nil {
		c.Error(fmt.Errorf("error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	customer, err := uc.customerService.GetBy("email", customerData.Email)
	if err != nil {
		c.Error(fmt.Errorf("error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if customer == nil {
		c.String(http.StatusNotFound, "User with the provided email was not found.")
		return
	}

	claims := jwt_utils.CreateClaims()
	claims.Email = customer.Email
	claims.UserId = customer.Id
	claims.Role = customer.Role

	tokenString, err := jwt_utils.CreateToken(claims)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (uc *CustomerController) SendRegistrationOAuthRequest(c *gin.Context) {
	reqURL := oauth_utils.GetRegistrationOAuthURL()
	c.Redirect(http.StatusFound, reqURL)
}

func (uc *CustomerController) OAuthRegistration(c *gin.Context) {

	code := c.Query("code")

	customerData, err := oauth_utils.GetCustomerGithubInformation(code, "REGISTRATION")
	if err != nil {
		c.Error(fmt.Errorf("error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	exists, err := uc.customerService.Exists(customerData.Email)
	if err != nil {
		c.Error(fmt.Errorf("error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		c.String(http.StatusNotFound, "User with the provided email already exists.")
		return
	}

	addrId, err := uc.addrService.CreateAddress(*customerData.Address)
	if err != nil {
		c.Error(fmt.Errorf("error while creating an address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	customerData.Address.Id = addrId

	if !exists {
		err = uc.customerService.CreateCustomer(*customerData)
		if err != nil {
			c.Error(fmt.Errorf("error while creating a customer. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Status(http.StatusCreated)
}

func (cc *CustomerController) GetCustomerInfo(c *gin.Context) {
	id := c.GetString("user_id")

	customerOut, err := cc.customerService.GetCustomerInfo(id)
	if err != nil {
		c.Error(fmt.Errorf("error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if customerOut == nil {
		c.String(http.StatusNotFound, "Customer doesn't exist.")
		return
	}

	c.JSON(http.StatusOK, customerOut)
}

// CUSTOMER UPDATE
func (cc *CustomerController) UpdateCustomer(c *gin.Context) {
	var customerID string = c.Param("id")
	var customerDto dto.CustomerInputDto
	err := c.BindJSON(&customerDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = cc.customerService.ValidateCustomerDataInput(customerDto)
	if err != nil {
		c.Error(fmt.Errorf("error while validating input. \nReason: %s", err.Error()))
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var addrId int
	addr, err := cc.addrService.GetAddr(*customerDto.Address)
	if err != nil {
		c.Error(fmt.Errorf("error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if addr == nil {
		addrId, err = cc.addrService.CreateAddress(*customerDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		customerDto.Address.Id = addrId
	} else {
		customerDto.Address.Id = addr.Id
	}

	res, err := cc.customerService.UpdateCustomer(customerID, &customerDto)
	if err != nil {
		c.Error(fmt.Errorf("error while updating the customer. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if !res {
		c.Error(fmt.Errorf("customer has failed to update or doesn't exist"))
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// TODO: FIX CUSTOMER PROFILE UPDATE
// TODO: CHECK FOR VALIDATIONS IN OTHER COMPONENTS
// CUSTOMER UPDATE

// PATCH http://localhost:3002/customer?attr=[NAME || SURNAME || AGE || USERNAME || EMAIL || PASSWORD ||ADDRESS]
// body: {value: [NAME || SURNAME || AGE || USERNAME || EMAIL || PASSWORD || ADDRESS]}

func (cc *CustomerController) UpdateCustomerProperty(c *gin.Context) {
	id := c.GetString("user_id")

	var property models.Property
	err := c.BindJSON(&property)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if len(property.Name) <= 0 || len(property.Value) <= 0 {
		c.Error(fmt.Errorf("invalid request"))
		c.Status(http.StatusBadRequest)
		return
	}

	err = cc.customerService.UpdateProperty(property.Name, property.Value, id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusAccepted, fmt.Sprintf("%s updated", property.Name))
}

func (cc *CustomerController) UpdateCustomerPassword(c *gin.Context) {
	id := c.GetString("user_id")

	var password models.Property
	err := c.BindJSON(&password)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = validations.ValidatePassword(password.Value)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	hash, err := security.HashPassword(password.Value)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err = cc.customerService.UpdateProperty("password", hash, id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusAccepted, "password updated")
}

func (cc *CustomerController) UpdateCustomerAddress(c *gin.Context) {
	userID := c.GetString("user_id")
	var address *dto.AddressInputDto
	err := c.BindJSON(&address)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = cc.addrService.ValidateAddress(address.City, address.Postfix, address.Street, address.StreetNum)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	addr, err := cc.addrService.GetAddr(*address)
	if err != nil {
		c.Error(fmt.Errorf("error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var addrID int
	if addr == nil {
		addrID, err = cc.addrService.CreateAddress(*address)
		if err != nil {
			c.Error(fmt.Errorf("error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	err = cc.customerService.UpdateProperty("address", addrID, userID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusAccepted, "address updated")
}
