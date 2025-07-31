package database

import (
	"carrierCheck/internal/domain/carrier"

	"gorm.io/gorm"
)

type CarrierRepository struct{
	Db *gorm.DB
}

func (c *CarrierRepository) Save(carrier *carrier.Carrier) error {
	tx := c.Db.Create(carrier)
	return tx.Error
}

func (c *CarrierRepository) GetAll() ([]carrier.Carrier, error) {
	var carriers []carrier.Carrier
	tx := c.Db.Preload("Email").Find(&carriers)
	return carriers, tx.Error
}

func (c *CarrierRepository) GetById(id string) (*carrier.Carrier, error) {
	var carrier carrier.Carrier
	tx := c.Db.Preload("Email").First(&carrier, "id = ?", id)
	return &carrier, tx.Error
}

func (c *CarrierRepository) Update(carrier *carrier.Carrier) error {
	tx := c.Db.Save(carrier)
	return tx.Error
}

func (c *CarrierRepository) Delete(carrier *carrier.Carrier) error {
	tx := c.Db.Delete(carrier)
	return tx.Error
}