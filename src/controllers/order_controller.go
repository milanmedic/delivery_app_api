package controllers

import (
	"fmt"
	"net/http"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	"delivery_app_api.mmedic.com/m/v2/src/services/order_service"
	"delivery_app_api.mmedic.com/m/v2/src/utils/validations"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	addrService  addr_service.AddrServicer
	orderService order_service.OrderServicer
}

func CreateOrderController(os order_service.OrderServicer, as addr_service.AddrServicer) *OrderController {
	return &OrderController{orderService: os, addrService: as}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var orderDto dto.OrderInputDto
	err := c.BindJSON(&orderDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	//TODO: Validate input
	//TODO: Refactor all address adding logic
	var addrId int
	addr, err := oc.addrService.GetAddr(orderDto.Address)
	valid := validations.ValidateAddress(*addr)
	if !valid {
		c.Error(fmt.Errorf("Error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err != nil {
		c.Error(fmt.Errorf("Error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if addr == nil {
		addrId, err = oc.addrService.CreateAddress(orderDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		orderDto.Address.Id = addrId
	} else {
		orderDto.Address.Id = addr.Id
	}

	err = oc.orderService.CreateOrder(orderDto)
	if err != nil {
		c.Error(fmt.Errorf("Error placing an order. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
	return
}

func (oc *OrderController) GetOrders(c *gin.Context) {
	id, ok := c.Get("user_id")

	if !ok {
		c.Error(fmt.Errorf("User not provided in token."))
		c.Status(http.StatusInternalServerError)
		return
	}

	orders, err := oc.orderService.GetOrdersByUserId(id.(string))
	if err != nil {
		c.Error(fmt.Errorf("Error retrieving orders. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if len(orders) <= 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(200, orders)
}
