package products

import (
	contracts "carrierCheck/internal/contracts/products"
	internalerrors "carrierCheck/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createProduct = contracts.CreateProduct{
		Name:  "Test Product",
		Price: 1.50,
		Qtde:  10,
	}
	service = ProductsService{}
)

type MockProductsRepository struct {
	mock.Mock
}

func (r *MockProductsRepository) Save(product *Product) error {
	args := r.Called(product)
	return args.Error(0)
}

func Test_Create_Products(t *testing.T) {
	assert := assert.New(t)
	mockProductsRepository := new(MockProductsRepository)
	mockProductsRepository.On("Save", mock.Anything).Return(nil)
	service.Repository = mockProductsRepository

	id, err := service.Create(createProduct)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveProducts(t *testing.T) {
	mockProductsRepository := new(MockProductsRepository)
	mockProductsRepository.On("Save", mock.MatchedBy(func(product *Product) bool {
		if (product.Name != createProduct.Name) ||
			(product.Price != createProduct.Price) ||
			(product.Qtde != createProduct.Qtde) {
			return false
		}

		return true
	})).Return(nil)

	service.Repository = mockProductsRepository
	service.Create(createProduct)

	mockProductsRepository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySaveCarrier(t *testing.T) {
	assert := assert.New(t)

	mockProductsRepository := new(MockProductsRepository)
	mockProductsRepository.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = mockProductsRepository
	_, err := service.Create(createProduct)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}