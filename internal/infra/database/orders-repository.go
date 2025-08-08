package database

import (
	"carrierCheck/internal/domain/carrier"
	"carrierCheck/internal/domain/clients"
	"carrierCheck/internal/domain/orders"
	"carrierCheck/internal/domain/products"

	"gorm.io/gorm"
)

type OrdersRepository struct{
	Db *gorm.DB
}

func (c *OrdersRepository) Save(order *orders.Order) error {
	tx := c.Db.Create(order)
	return tx.Error
}

func (c *OrdersRepository) GetAll() ([]orders.Order, error) {
	var orders []orders.Order
	tx := c.Db.Preload("OrderProduct").Find(&orders)
	return orders, tx.Error
}

func (c *OrdersRepository) GetById(id string) (*orders.Order, error) {
	var order orders.Order
	tx := c.Db.Preload("OrderProduct").First(&order, "id = ?", id)
	return &order, tx.Error
}

func (c *OrdersRepository) Update(order *orders.Order) error {
	tx := c.Db.Save(order)
    return tx.Error
}

func (c *OrdersRepository) Delete(order *orders.Order) error {
	tx := c.Db.Delete(order)
	return tx.Error
}

func (c *OrdersRepository) GetByIdClients(id string) (*clients.Client, error) {
	var client clients.Client
	tx := c.Db.Preload("Address").First(&client, "id = ?", id)
	return &client, tx.Error
}

func (c *OrdersRepository) GetByIdAddress(id string) (*clients.AddressClients, error) {
	var address clients.AddressClients
	tx := c.Db.First(&address, "id = ?", id)
	return &address, tx.Error
}

func (c *OrdersRepository) GetByIdCarrier(id string) (*carrier.Carrier, error) {
	var carrier carrier.Carrier
	tx := c.Db.Preload("Email").First(&carrier, "id = ?", id)
	return &carrier, tx.Error
}

func (c *OrdersRepository) GetByIdProduct(id string) (*products.Product, error) {
	var product products.Product
	tx := c.Db.First(&product, "id = ?", id)
	return &product, tx.Error
}