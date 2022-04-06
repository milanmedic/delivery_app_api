package controllers

import (
	addr_service "delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	"delivery_app_api.mmedic.com/m/v2/src/services/deliverer_service"
)

type DelivererController struct {
	delivererService deliverer_service.DelivererServicer
	addrService      addr_service.AddrServicer
}

func CreateDelivererController(ds deliverer_service.DelivererServicer, as addr_service.AddrServicer) *DelivererController {
	return &DelivererController{delivererService: ds, addrService: as}
}
