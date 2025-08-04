package database

import (
	"carrierCheck/internal/domain/clients"

	"gorm.io/gorm"
)

type ClientsRepository struct{
	Db *gorm.DB
}

func (c *ClientsRepository) Save(client *clients.Client) error {
	tx := c.Db.Create(client)
	return tx.Error
}

func (c *ClientsRepository) GetAll() ([]clients.Client, error) {
	var clients []clients.Client
	tx := c.Db.Preload("Address").Find(&clients)
	return clients, tx.Error
}

func (c *ClientsRepository) GetById(id string) (*clients.Client, error) {
	var client clients.Client
	tx := c.Db.Preload("Address").First(&client, "id = ?", id)
	return &client, tx.Error
}

func (c *ClientsRepository) Update(client *clients.Client) error {
	tx := c.Db.Save(client)
    return tx.Error
}

func (c *ClientsRepository) Delete(client *clients.Client) error {
	tx := c.Db.Delete(client)
	return tx.Error
}