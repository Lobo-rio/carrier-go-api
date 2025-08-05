package endpoints

import (
	"carrierCheck/internal/domain/carrier"
	"carrierCheck/internal/domain/clients"
	"carrierCheck/internal/domain/orders"
	"carrierCheck/internal/domain/products"
)

type Handler struct {
	CarrierService carrier.CarrierService
	ClientsService clients.ClientsService
	ProductsService products.ProductsService
	OrdersService orders.OrdersService
}