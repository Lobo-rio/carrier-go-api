package orders

import (
	"carrierCheck/internal/domain/carrier"
	"carrierCheck/internal/domain/clients"
	"carrierCheck/internal/domain/products"
)

type OrdersRepository interface {
	Save(order *Order) error
	GetAll() ([]Order, error)
	GetById(id string) (*Order, error)
	Update(order *Order) error
	Delete(order *Order) error
	GetByIdClients(id string) (*clients.Client, error)
	GetByIdAddress(id string) (*clients.AddressClients, error)
	GetByIdCarrier(id string) (*carrier.Carrier, error)
	GetByIdProduct(id string) (*products.Product, error)
}
