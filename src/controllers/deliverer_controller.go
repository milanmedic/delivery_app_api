package controllers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	addr_service "delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	"delivery_app_api.mmedic.com/m/v2/src/services/deliverer_service"
	"delivery_app_api.mmedic.com/m/v2/src/services/order_service"
	"delivery_app_api.mmedic.com/m/v2/src/utils/jwt_utils"
	"delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"delivery_app_api.mmedic.com/m/v2/src/utils/validations"
	"github.com/gin-gonic/gin"
)

type DelivererController struct {
	delivererService deliverer_service.DelivererServicer
	addrService      addr_service.AddrServicer
	orderService     order_service.OrderServicer
}

func CreateDelivererController(ds deliverer_service.DelivererServicer, as addr_service.AddrServicer, os order_service.OrderServicer) *DelivererController {
	return &DelivererController{delivererService: ds, addrService: as, orderService: os}
}

func (dc *DelivererController) Register(c *gin.Context) {
	var delivererDto dto.DelivererInputDto
	err := c.BindJSON(&delivererDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = dc.delivererService.ValidateDelivererDataInput(delivererDto)
	if err != nil {
		c.Error(fmt.Errorf("Error while validating input. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var addrId int
	addr, err := dc.addrService.GetAddr(*delivererDto.Address)
	if err != nil {
		c.Error(fmt.Errorf("Error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if addr == nil {
		addrId, err = dc.addrService.CreateAddress(*delivererDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		delivererDto.Address.Id = addrId
	} else {
		delivererDto.Address.Id = addr.Id
	}

	exists, err := dc.delivererService.Exists(delivererDto.Email)
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
		err = dc.delivererService.AddDeliverer(delivererDto)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating a customer. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusCreated, "Created.")
	return
}

func (dc *DelivererController) DelivererLogin(c *gin.Context) {
	var credentials jwt_utils.Credentials
	err := c.BindJSON(&credentials)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = validations.ValidateEmail(credentials.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	deliverer, err := dc.delivererService.GetBy("email", credentials.Email)
	if err != nil {
		c.Error(fmt.Errorf("error while retrieving the deliverer info. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if deliverer == nil {
		c.JSON(http.StatusNotFound, "Deliverer with the provided email was not found.")
		return
	}

	if !security.CheckPasswordHash(credentials.Password, deliverer.Password) {
		c.JSON(http.StatusUnauthorized, "Wrong password.")
		return
	}

	if strings.Compare(deliverer.VerificationStatus, "VERIFIED") != 0 {
		c.JSON(http.StatusUnauthorized, "Account not verified.")
		return
	}

	claims := jwt_utils.CreateClaims()
	claims.Email = credentials.Email
	claims.UserId = deliverer.Id
	claims.Role = deliverer.Role

	tokenString, err := jwt_utils.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (dc *DelivererController) GetDelivererInfo(c *gin.Context) {
	var id string = c.GetString("user_id")

	delivererOut, err := dc.delivererService.GetDelivererInfo(id)
	if err != nil {
		c.Error(fmt.Errorf("error while retrieving the deliverer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if delivererOut == nil {
		c.String(http.StatusNotFound, "Deliverer doesn't exist.")
		return
	}

	c.JSON(http.StatusOK, delivererOut)
}

func (dc *DelivererController) UpdateDeliverer(c *gin.Context) {
	var delivererID string = c.Param("id")
	var delivererDto dto.DelivererInputDto
	err := c.BindJSON(&delivererDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = dc.delivererService.ValidateDelivererDataInput(delivererDto)
	if err != nil {
		c.Error(fmt.Errorf("error while validating input. \nReason: %s", err.Error()))
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var addrId int
	addr, err := dc.addrService.GetAddr(*delivererDto.Address)
	if err != nil {
		c.Error(fmt.Errorf("error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if addr == nil {
		addrId, err = dc.addrService.CreateAddress(*delivererDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		delivererDto.Address.Id = addrId
	} else {
		delivererDto.Address.Id = addr.Id
	}

	res, err := dc.delivererService.UpdateDeliverer(delivererID, &delivererDto)
	if err != nil {
		c.Error(fmt.Errorf("error while updating the deliverer. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if !res {
		c.Error(fmt.Errorf("deliverer has failed to update or doesn't exist"))
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (dc *DelivererController) AcceptOrder(c *gin.Context) {
	var orderID string = c.Param("id")
	delivererID := c.GetString("user_id")

	deliverer, err := dc.delivererService.GetBy("id", delivererID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if deliverer.DeliveryInProgress {
		c.Error(fmt.Errorf("you already have a delivery in progress"))
		c.JSON(http.StatusBadRequest, "you already have a delivery in progress")
		return
	}

	// TODO:
	// AFTER RANDOM TIME FROM ORDER ACCEPTANCE
	// CREATE A NEW ROUTINE TO UPDATE ORDER DELIVERY TIME
	// UPDATE DELIVERY TIME

	// CREATE A NEW ROUTINE FROM MAIN
	// UPDATE ALL ROUTES WHERE ORDER_ACCEPT_TIME == ORDER_DELIVERY_TIME
	order, err := dc.orderService.GetOrder(orderID)
	if err != nil {
		c.Error(fmt.Errorf("error while accepting the order. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if order.Accepted {
		c.JSON(http.StatusForbidden, "Order already taken.")
		return
	}

	err = dc.orderService.AcceptOrder(orderID, delivererID)
	if err != nil {
		c.Error(fmt.Errorf("error while accepting the order. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = dc.delivererService.UpdateProperty("delivery_in_progress", true, delivererID)
	if err != nil {
		c.Error(fmt.Errorf("error while accepting the order. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = dc.orderService.UpdateProperty("status", "IN_PROGRESS", orderID)
	if err != nil {
		c.Error(fmt.Errorf("error while accepting the order. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, fmt.Sprintf("Order %s accepted", orderID))
}

func (dc *DelivererController) StreamDeliveryTime(c *gin.Context) {
	// GET USER ID FROM TOKEN
	// POOL DB
	c.Stream(func(w io.Writer) bool {
		c.SSEvent("delivery_time", "hello")
		c.Done()
		return true
	})
}

// PROFILE UPDATE

func (dc *DelivererController) UpdateDelivererProperty(c *gin.Context) {
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

	err = dc.delivererService.UpdateProperty(property.Name, property.Value, id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusAccepted, fmt.Sprintf("%s updated", property.Name))
}

func (dc *DelivererController) UpdateDelivererPassword(c *gin.Context) {
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

	err = dc.delivererService.UpdateProperty("password", hash, id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusAccepted, "password updated")
}

func (dc *DelivererController) UpdateDelivererAddress(c *gin.Context) {
	userID := c.GetString("user_id")
	var address *dto.AddressInputDto
	err := c.BindJSON(&address)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = dc.addrService.ValidateAddress(address.City, address.Postfix, address.Street, address.StreetNum)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	addr, err := dc.addrService.GetAddr(*address)
	if err != nil {
		c.Error(fmt.Errorf("error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var addrID int
	if addr == nil {
		addrID, err = dc.addrService.CreateAddress(*address)
		if err != nil {
			c.Error(fmt.Errorf("error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	err = dc.delivererService.UpdateProperty("address", addrID, userID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusAccepted, "address updated")
}

func (dc *DelivererController) CompleteOrder(c *gin.Context) {
	orderID := c.Param("id")
	delivererID := c.GetString("user_id")

	status, err := dc.orderService.GetOrderStatus(orderID)
	if err != nil {
		c.Error(fmt.Errorf("order completion failed. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if strings.Compare(status, "COMPLETED") == 0 || strings.Compare(status, "PENDING") == 0 || strings.Compare(status, "CANCELLED") == 0 {
		c.Error(fmt.Errorf("order completion failed. Cannot cancel order that is in progress or completed"))
		c.JSON(http.StatusBadRequest, "Order completion failed. Cannot complete order that is in progress, completed or cancelled.")
		return
	}

	err = dc.orderService.UpdateProperty("status", "COMPLETED", orderID)
	if err != nil {
		c.Error(fmt.Errorf("order completion failed. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = dc.delivererService.UpdateProperty("delivery_in_progress", false, delivererID)
	if err != nil {
		c.Error(fmt.Errorf("order completion failed. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// if status noContent present, there is no body in req
	c.JSON(http.StatusOK, "Completed")
}
