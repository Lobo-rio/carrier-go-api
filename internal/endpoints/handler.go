package endpoints

import (
	"carrierCheck/internal/domain/carrier"
	"carrierCheck/internal/domain/clients"
)

type Handler struct {
	CarrierService carrier.CarrierService
	ClientsService clients.ClientsService
}