package carrier

import (
	contracts "carrierCheck/internal/contracts/carrier"
	internalerrors "carrierCheck/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createCarrier = contracts.CreateCarrier{
		Name:    "Test Carrier",
		Email:   []string{"test@example.com", "test2@example.com"},
		Phone:   "2134567890",
		Contact: "John Doe",
	}
	service = CarrierService{}
)

type MockCarrierRepository struct {
	mock.Mock
}

func (r *MockCarrierRepository) Save(carrier *Carrier) error {
	args := r.Called(carrier)
	return args.Error(0)
}

func (r *MockCarrierRepository) GetAll() ([]Carrier, error) {
	args := r.Called()
	return args.Get(0).([]Carrier), args.Error(1)
}

func Test_Create_Carrier(t *testing.T) {
	assert := assert.New(t)
	mockCarrierRepository := new(MockCarrierRepository)
	mockCarrierRepository.On("Save", mock.Anything).Return(nil)
	service.Repository = mockCarrierRepository

	id, err := service.Create(createCarrier)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCarrier(t *testing.T) {
	mockCarrierRepository := new(MockCarrierRepository)
    mockCarrierRepository.On("Save", mock.MatchedBy(func (carrier *Carrier) bool {
		if (carrier.Name != createCarrier.Name) || 
		   (carrier.Phone != createCarrier.Phone) || 
		   (carrier.Contact != createCarrier.Contact) || 
		   (len(carrier.Email) != len(createCarrier.Email)) { 
			return false 
		}

		return true
	})).Return(nil)

	service.Repository = mockCarrierRepository
	service.Create(createCarrier)

	mockCarrierRepository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySaveCarrier(t *testing.T) {
	assert := assert.New(t)
	
	mockCarrierRepository := new(MockCarrierRepository)
    mockCarrierRepository.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = mockCarrierRepository
	_, err := service.Create(createCarrier)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}