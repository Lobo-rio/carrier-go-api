package carrier

import (
	contracts "carrierCheck/internal/contracts/carrier"
	internalerrors "carrierCheck/internal/internal-errors"
)

type CarrierService struct {
	Repository CarrierRepository
}

func (s *CarrierService) Create(createCarrier contracts.CreateCarrier) (string, error) {
	carrier, err := NewCarrier(createCarrier.Name, createCarrier.Phone, createCarrier.Contact, createCarrier.Email)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(carrier)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return carrier.ID, nil
}