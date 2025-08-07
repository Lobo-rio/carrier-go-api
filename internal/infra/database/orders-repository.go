package database

import (
	"carrierCheck/internal/domain/orders"

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
	var models []orders.Order
	tx := c.Db.Find(&models)
	return models, tx.Error
}

func (c *OrdersRepository) GetById(id string) (*orders.Order, error) {
	var order orders.Order
	tx := c.Db.First(&order, "id = ?", id)
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