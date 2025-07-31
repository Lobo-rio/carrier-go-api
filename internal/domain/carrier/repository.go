package carrier

import contracts "carrierCheck/internal/contracts/carrier"

type CarrierRepository interface {
	Save(carrier *Carrier) error
	GetAll() ([]Carrier, error)
	GetById(id string) (*Carrier, error)
	Update(id string, request *contracts.UpdateCarrier) error
	Delete(carrier *Carrier) error
}