package clients

import (
	contracts "carrierCheck/internal/contracts/clients"
	internalerrors "carrierCheck/internal/internal-errors"
)

type ClientsService struct {
	Repository ClientsRepository
}

func (s *ClientsService) Create(createClient contracts.CreateClient) (string, error) {
	client, err := NewClient(createClient.Name, createClient.Email, createClient.Phone)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(client)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return client.ID, nil
}