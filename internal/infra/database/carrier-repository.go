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
	tx := c.Db.Find(&carriers)
	return carriers, tx.Error
}

func (c *CarrierRepository) GetByID(id string) (*carrier.Carrier, error) {
	var carrier carrier.Carrier
	tx := c.Db.First(&carrier, id)
	return &carrier, tx.Error
}