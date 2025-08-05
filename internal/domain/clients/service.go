package clients

import (
	contracts "carrierCheck/internal/contracts/clients"
	internalerrors "carrierCheck/internal/internal-errors"
)

type ClientsService interface {
	Create(createClient contracts.CreateClient) (string, error)
	GetById(id string) (*contracts.ResponseClient, error)
	GetAll() ([]contracts.ResponseClient, error)
	Update(id string, request contracts.UpdateClient) error
	Delete(id string)  error
}

type ClientsServiceImp struct {
	Repository ClientsRepository
}

func (s *ClientsServiceImp) Create(createClient contracts.CreateClient) (string, error) {
	addressClients := make([]AddressClients, len(createClient.Address))
	for i, addr := range createClient.Address {
		addressClients[i] = AddressClients{
			Address:      addr.Address,
			Number:      addr.Number,
			Complement:  addr.Complement,
			Neighborhood: addr.Neighborhood,
			City:        addr.City,
			State:       addr.State,
		}
	}
	client, err := NewClient(createClient.Name, createClient.Email, createClient.Phone, addressClients)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(client)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return client.ID, nil
}

func (s *ClientsServiceImp) GetById(id string) (*contracts.ResponseClient, error) {
	client, err := s.Repository.GetById(id)
	if err != nil {
		return nil, internalerrors.ErrInternal
	}
	addressClient := make([]AddressClients, len(client.Address))
	for i, address := range client.Address {
		addressClient[i].Address = address.Address 
		addressClient[i].Number = address.Number 
		addressClient[i].Complement = address.Complement 
		addressClient[i].Neighborhood = address.Neighborhood 
		addressClient[i].City = address.City 
		addressClient[i].State = address.State 
	}
	return &contracts.ResponseClient{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		Address:   addressClient,
		CreatedAt: client.CreatedAt.String(),
	}, nil
}

func (s *ClientsServiceImp) GetAll() ([]contracts.ResponseClient, error) {
	clients, err := s.Repository.GetAll()
	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	responseClients := make([]contracts.ResponseClient, len(clients))
	for i, client := range clients {
		addressClient := make([]AddressClients, len(client.Address))
		for j, address := range client.Address {
			addressClient[j].Address = address.Address
			addressClient[j].Number = address.Number
			addressClient[j].Complement = address.Complement
			addressClient[j].Neighborhood = address.Neighborhood
			addressClient[j].City = address.City
			addressClient[j].State = address.State
		}
		responseClients[i] = contracts.ResponseClient{
			ID:        client.ID,
			Name:      client.Name,
			Email:     client.Email,
			Phone:     client.Phone,
			Address:   addressClient,
			CreatedAt: client.CreatedAt.String(),
		}
	}

	return responseClients, nil
}

func (s *ClientsServiceImp) Update(id string, request contracts.UpdateClient) error {
	client, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ErrInternal
	}

	client.Name = request.Name
	client.Email = request.Email
	client.Phone = request.Phone

	err = s.Repository.Update(client)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}

func (s *ClientsServiceImp) Delete(id string)  error {
	carrier, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ErrInternal
	}

	err = s.Repository.Delete(carrier)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}