package orders

import (
	"testing"

	contracts "carrierCheck/internal/contracts/order"
	carrierDomain "carrierCheck/internal/domain/carrier"
	clientsDomain "carrierCheck/internal/domain/clients"
	productsDomain "carrierCheck/internal/domain/products"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrdersRepository struct {
	mock.Mock
}

func (m *MockOrdersRepository) Save(order *Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrdersRepository) GetById(id string) (*Order, error) {
	args := m.Called(id)
	return args.Get(0).(*Order), args.Error(1)
}

func (m *MockOrdersRepository) GetAll() ([]Order, error) {
	args := m.Called()
	return args.Get(0).([]Order), args.Error(1)
}

func (m *MockOrdersRepository) Update(order *Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrdersRepository) Delete(order *Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrdersRepository) GetByIdClients(id string) (*clientsDomain.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*clientsDomain.Client), args.Error(1)
}

func (m *MockOrdersRepository) GetByIdAddress(id string) (*clientsDomain.AddressClients, error) {
	args := m.Called(id)
	return args.Get(0).(*clientsDomain.AddressClients), args.Error(1)
}

func (m *MockOrdersRepository) GetByIdCarrier(id string) (*carrierDomain.Carrier, error) {
	args := m.Called(id)
	return args.Get(0).(*carrierDomain.Carrier), args.Error(1)
}

func (m *MockOrdersRepository) GetByIdProduct(id string) (*productsDomain.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*productsDomain.Product), args.Error(1)
}

const OrderShipped = "Pedido Enviado"
const ErrNotFound = "not found"

func TestOrdersService_Create(t *testing.T) {
	repo := new(MockOrdersRepository)
	svc := &OrdersServiceImp{
		Repository: repo,
		SendMail: func(to string, subject string, body string) error {
			return nil
		},
	}

	createOrder := contracts.CreateOrder{
		ClientId:  "client1",
		AddressId: "address1",
		Products: []contracts.OrderProduct{
			{
				ProductId: "product1",
				Quantity:  2,
				Price:     100.0,
			},
			{
				ProductId: "product2",
				Quantity:  1,
				Price:     50.0,
			},
		},
	}

	repo.On("Save", mock.Anything).Return(nil)

	id, err := svc.Create(createOrder)

	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	repo.AssertCalled(t, "Save", mock.Anything)
}

func TestOrdersService_GetById(t *testing.T) {
	repo := new(MockOrdersRepository)
	svc := &OrdersServiceImp{
		Repository: repo,
	}

	order := &Order{
		ID:        "order1",
		ClientId:  "client1",
		AddressId: "address1",
		Status:    OrderPlaced,
		OrderProduct: []OrdersProducts{
			{
				ProductId: "product1",
				Quantity:  2,
				Price:     100.0,
			},
			{
				ProductId: "product2",
				Quantity:  1,
				Price:     50.0,
			},
		},
	}

	repo.On("GetById", "order1").Return(order, nil)

	response, err := svc.GetById("order1")

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "order1", response.ID)
	assert.Equal(t, "client1", response.ClientId)
	assert.Equal(t, "address1", response.AddressId)
	assert.Len(t, response.Products, 2)
	repo.AssertCalled(t, "GetById", "order1")
}

func TestOrdersService_Update(t *testing.T) {
	repo := new(MockOrdersRepository)
	svc := &OrdersServiceImp{
		Repository: repo,
	}

	order := &Order{
		ID:        "order1",
		ClientId:  "client1",
		AddressId: "address1",
		Status:    OrderPlaced,
	}

	repo.On("GetById", "order1").Return(order, nil)
	repo.On("Update", mock.Anything).Return(nil)

	updateRequest := contracts.UpdateOrder{
		ClientId:  "client2",
		AddressId: "address2",
	}

	err := svc.Update("order1", updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, "client2", order.ClientId)
	assert.Equal(t, "address2", order.AddressId)
	repo.AssertCalled(t, "GetById", "order1")
	repo.AssertCalled(t, "Update", mock.Anything)
}

func TestOrdersService_Delete(t *testing.T) {
	repo := new(MockOrdersRepository)
	svc := &OrdersServiceImp{
		Repository: repo,
	}

	order := &Order{
		ID:        "order1",
		ClientId:  "client1",
		AddressId: "address1",
		Status:    OrderPlaced,
	}

	repo.On("GetById", "order1").Return(order, nil)
	repo.On("Delete", order).Return(nil)

	err := svc.Delete("order1")

	assert.NoError(t, err)
	repo.AssertCalled(t, "GetById", "order1")
	repo.AssertCalled(t, "Delete", order)
}

func TestOrdersService_GetAll(t *testing.T) {
	repo := new(MockOrdersRepository)
	svc := &OrdersServiceImp{
		Repository: repo,
	}

	orders := []Order{
		{
			ID:        "order1",
			ClientId:  "client1",
			AddressId: "address1",
			Status:    OrderPlaced,
			OrderProduct: []OrdersProducts{
				{
					ProductId: "product1",
					Quantity:  2,
					Price:     100.0,
				},
				{
					ProductId: "product2",
					Quantity:  1,
					Price:     50.0,
				},
			},
		},
		{
			ID:        "order2",
			ClientId:  "client2",
			AddressId: "address2",
			Status:    OrderShipped,
			OrderProduct: []OrdersProducts{
				{
					ProductId: "product3",
					Quantity:  1,
					Price:     150.0,
				},
			},
		},
	}

	repo.On("GetAll").Return(orders, nil)

	response, err := svc.GetAll()

	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, "order1", response[0].ID)
	assert.Equal(t, "client1", response[0].ClientId)
	assert.Equal(t, "address1", response[0].AddressId)
	assert.Len(t, response[0].Products, 2)
	assert.Equal(t, "order2", response[1].ID)
	assert.Equal(t, "client2", response[1].ClientId)
	assert.Equal(t, "address2", response[1].AddressId)
	assert.Len(t, response[1].Products, 1)
	repo.AssertCalled(t, "GetAll")
}

func TestOrdersService_Create_Invalid(t *testing.T) {
	repo := new(MockOrdersRepository)
	svc := &OrdersServiceImp{
		Repository: repo,
	}

	invalidOrder := contracts.CreateOrder{
		ClientId:  "",
		AddressId: "address1",
		Products: []contracts.OrderProduct{
			{
				ProductId: "product1",
				Quantity:  0,
				Price:     100.0,
			},
		},
	}

	id, err := svc.Create(invalidOrder)

	assert.Error(t, err)
	assert.Empty(t, id)
}

func TestOrdersService_Update_NonExistentOrder(t *testing.T) {
	repo := new(MockOrdersRepository)
	svc := &OrdersServiceImp{
		Repository: repo,
	}

	repo.On("GetById", "order1").Return(nil, ErrNotFound)

	updateRequest := contracts.UpdateOrder{
		ClientId:  "client2",
		AddressId: "address2",
	}

	err := svc.Update("order1", updateRequest)

	assert.Error(t, err)
	repo.AssertCalled(t, "GetById", "order1")
}

func TestOrdersService_Delete_NonExistentOrder(t *testing.T) {
	repo := new(MockOrdersRepository)
	svc := &OrdersServiceImp{
		Repository: repo,
	}

	repo.On("GetById", "order1").Return(nil, ErrNotFound)

	err := svc.Delete("order1")

	assert.Error(t, err)
	repo.AssertCalled(t, "GetById", "order1")
}