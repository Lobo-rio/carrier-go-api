package carrier

import (
	contracts "carrierCheck/internal/contracts/carrier"
	internalerrors "carrierCheck/internal/internal-errors"
)

type CarrierService interface {
	Create(createCarrier contracts.CreateCarrier) (string, error)
	GetById(id string) (*contracts.ResponseCarrier, error)
	GetAll() ([]contracts.ResponseCarrier, error)
	Update(id string, request contracts.UpdateCarrier) error
	Delete(id string)  error
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

func (s *CarrierServiceImp) GetById(id string) (*contracts.ResponseCarrier, error) {
	carrier, err := s.Repository.GetById(id)
	if err != nil {
		return nil, internalerrors.ProcessErrorToReturn(err)
	}
	emails := make([]string, len(carrier.Email))
	for i, e := range carrier.Email {
		emails[i] = e.Email 
	}
	return &contracts.ResponseCarrier{
		ID:        carrier.ID,
		Name:      carrier.Name,
		Email:     emails,
		Phone:     carrier.Phone,
		Contact:   carrier.Contact,
		CreatedAt: carrier.CreatedAt.String(),
	}, nil
}

func (s *CarrierServiceImp) GetAll() ([]contracts.ResponseCarrier, error) {
	carriers, err := s.Repository.GetAll()
	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	responseCarriers := make([]contracts.ResponseCarrier, len(carriers))
	for i, carrier := range carriers {
		emails := make([]string, len(carrier.Email))
		for j, e := range carrier.Email {
			emails[j] = e.Email
		}
		responseCarriers[i] = contracts.ResponseCarrier{
			ID:        carrier.ID,
			Name:      carrier.Name,
			Email:     emails,
			Phone:     carrier.Phone,
			Contact:   carrier.Contact,
			CreatedAt: carrier.CreatedAt.String(),
		}
	}

	return responseCarriers, nil
}

func (s *CarrierServiceImp) Update(id string, request contracts.UpdateCarrier) error {
	carrier, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	carrier.Name = request.Name
	carrier.Phone = request.Phone
	carrier.Contact = request.Contact

	err = s.Repository.Update(carrier)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}

func (s *CarrierServiceImp) Delete(id string)  error {
	carrier, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	err = s.Repository.Delete(carrier)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}