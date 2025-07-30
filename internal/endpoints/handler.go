package endpoints

import "carrierCheck/internal/domain/carrier"

type Handler struct {
	CarrierService carrier.CarrierService
}