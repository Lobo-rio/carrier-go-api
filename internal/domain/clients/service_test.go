package clients

import (
	contracts "carrierCheck/internal/contracts/clients"
	internalerrors "carrierCheck/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createClient = contracts.CreateClient{
		Name:  "Test Client",
		Email: "test@example.com",
		Phone: "2134567890",
	}
	service = ClientsService{}
)

type MockClientsRepository struct {
	mock.Mock
}

func (r *MockClientsRepository) Save(client *Client) error {
	args := r.Called(client)
	return args.Error(0)
}

func Test_Create_Clients(t *testing.T) {
	assert := assert.New(t)
	mockClientsRepository := new(MockClientsRepository)
	mockClientsRepository.On("Save", mock.Anything).Return(nil)
	service.Repository = mockClientsRepository

	id, err := service.Create(createClient)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveClients(t *testing.T) {
	mockClientsRepository := new(MockClientsRepository)
    mockClientsRepository.On("Save", mock.MatchedBy(func (client *Client) bool {
		if (client.Name != createClient.Name) || 
		   (client.Phone != createClient.Phone) || 
		   (client.Email != createClient.Email) { 
			return false 
		}

		return true
	})).Return(nil)

	service.Repository = mockClientsRepository
	service.Create(createClient)

	mockClientsRepository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySaveCarrier(t *testing.T) {
	assert := assert.New(t)
	
	mockClientsRepository := new(MockClientsRepository)
    mockClientsRepository.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = mockClientsRepository
	_, err := service.Create(createClient)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}