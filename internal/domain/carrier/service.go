package carrier

import (
	contracts "carrierCheck/internal/contracts/carrier"
	internalerrors "carrierCheck/internal/internal-errors"
)

type CarrierService interface {
	Create(createCarrier contracts.CreateCarrier) (string, error)
}
type CarrierServiceImp struct {
	Repository CarrierRepository
}

func (s *CarrierServiceImp) Create(createCarrier contracts.CreateCarrier) (string, error) {
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