package carrier

import (
	contracts "carrierCheck/internal/contracts/carrier"
)

type CarrierService struct {
	Repository CarrierRepository
}

func (s *CarrierService) Create(createCarrier contracts.CreateCarrier) (string, error) {
	carrier, _ := NewCarrier(createCarrier.Name, createCarrier.Phone, createCarrier.Contact, createCarrier.Email)
	s.Repository.Save(carrier)

	return carrier.ID, nil
}