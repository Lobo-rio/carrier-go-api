package endpoints

import (
	"bytes"
	contracts "carrierCheck/internal/contracts/carrier"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	body = contracts.CreateCarrier{
		Name:    "Test Carrier",
		Email:   []string{"test@example.com", "test2@example.com"},
		Phone:   "2134567890",
		Contact: "John Doe",
	}
)

type MockCarrierService struct {
	mock.Mock
}

func (r *MockCarrierService) Create(createCarrier contracts.CreateCarrier) (string, error) {
	args := r.Called(createCarrier)
	return args.String(0), args.Error(1)
}

func Test_CreateCarrier_Save(t *testing.T) {
    assert := assert.New(t)
	service := new(MockCarrierService)
	service.On("Create", mock.MatchedBy(func(createCarrier contracts.CreateCarrier) bool {
		if createCarrier.Name != body.Name ||
			createCarrier.Phone != body.Phone ||
			createCarrier.Contact != body.Contact ||
			len(createCarrier.Email) != len(body.Email) {
			return false
		}
		return true
	})).Return("12345x", nil)
	handler := Handler{
		CarrierService: service,
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/carriers", &buf)
	rr := httptest.NewRecorder()
	
	_, status, err := handler.CreateCarrier(rr, req)

    assert.Equal(201, status)
	assert.Nil(err)
}

func Test_CreateCarrier_Return_Error_Save(t *testing.T) {
    assert := assert.New(t)
	service := new(MockCarrierService)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error saving carrier"))
	handler := Handler{
		CarrierService: service,
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/carriers", &buf)
	rr := httptest.NewRecorder()
	
	_, _, err := handler.CreateCarrier(rr, req)

    assert.NotNil(err)
}
