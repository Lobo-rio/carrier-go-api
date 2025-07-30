package database

import "carrierCheck/internal/domain/carrier"

type CarrierRepository struct{
	carriers []carrier.Carrier
}

func (c *CarrierRepository) Save(carrier *carrier.Carrier) error {
	c.carriers = append(c.carriers, *carrier)
	return nil
}

func (c *CarrierRepository) GetAll() ([]carrier.Carrier, error) {
	return c.carriers, nil
}