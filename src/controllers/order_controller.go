package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
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
	orderDto.CustomerID = c.GetString("user_id")

	err := c.BindJSON(&orderDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request.")
		return
	}

	//TODO: Validate input
	//TODO: Refactor all address adding logic
	var addrId int
	addr, err := oc.addrService.GetAddr(orderDto.Address)

	if addr == nil {
		valid := validations.ValidateAddress(orderDto.Address)
		if !valid {
			c.Error(fmt.Errorf("error while validating for address. \nReason: %s", err.Error()))
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	if err != nil {
		c.Error(fmt.Errorf("error while searching for address. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if addr == nil {
		addrId, err = oc.addrService.CreateAddress(orderDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("error while creating an address. \nReason: %s", err.Error()))
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		orderDto.Address.Id = addrId
	} else {
		orderDto.Address.Id = addr.Id
	}

	err = oc.orderService.CreateOrder(orderDto)
	if err != nil {
		c.Error(fmt.Errorf("error placing an order. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "Created.")
}

func (oc *OrderController) GetOrders(c *gin.Context) {
	id, ok := c.Get("user_id")

	if !ok {
		c.Error(fmt.Errorf("user not provided in token"))
		c.Status(http.StatusUnauthorized)
		return
	}

	orders, err := oc.orderService.GetOrdersByUserId(id.(string))
	if err != nil {
		c.Error(fmt.Errorf("error retrieving orders. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(orders) <= 0 {
		c.JSON(http.StatusNotFound, orders)
		return
	}

	c.JSON(200, orders)
}

func (oc *OrderController) CancelOrder(c *gin.Context) {
	id := c.Param("id")

	status, err := oc.orderService.GetOrderStatus(id)
	if err != nil {
		c.Error(fmt.Errorf("order cancellation failed. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if strings.Compare(status, "IN_PROGRESS") == 0 || strings.Compare(status, "COMPLETED") == 0 || strings.Compare(status, "CANCELLED") == 0 {
		c.Error(fmt.Errorf("order cancellation failed. Cannot cancel order that is in progress or completed"))
		c.JSON(http.StatusBadRequest, "Order cancellation failed. Cannot cancel order that is in progress, completed or cancelled.")
		return
	}

	err = oc.orderService.CancelOrder(id)
	if err != nil {
		c.Error(fmt.Errorf("order cancellation failed. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// if status noContent present, there is no body in req
	c.JSON(http.StatusOK, "Cancelled")
}

func (oc *OrderController) GetAllOrders(c *gin.Context) {
	deliveryStatus := c.Query("delivery_status")
	accepted := c.Query("accepted")

	var orders []models.Order
	var err error
	if strings.Compare(deliveryStatus, "") == 0 && strings.Compare(accepted, "") == 0 {
		orders, err = oc.orderService.GetAllOrders("")
	} else if strings.Compare(deliveryStatus, "") != 0 && strings.Compare(accepted, "") == 0 {
		orders, err = oc.orderService.GetAllOrders(deliveryStatus)
	} else {
		if strings.Compare(accepted, "true") == 0 {
			accepted = "1"
		} else {
			accepted = "0"
		}
		orders, err = oc.orderService.GetAllOrders(deliveryStatus, accepted)
	}

	if err != nil {
		c.Error(fmt.Errorf("error retrieving orders. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(orders) <= 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(200, orders)
}

func (oc *OrderController) GetAllDelivererOrders(c *gin.Context) {
	id := c.GetString("user_id")

	orders, err := oc.orderService.GetOrdersByDelivererId(id)
	if err != nil {
		c.Error(fmt.Errorf("error retrieving orders. \nReason: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(orders) <= 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(200, orders)
}
