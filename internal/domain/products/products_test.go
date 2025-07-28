package products

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
    name   = "Test Product"
    price  = 1.50
    qtde   = 10
)

func Test_NewProduct_Create(t *testing.T) {
    assert := assert.New(t)

	product, _ := NewProduct(name, price, qtde)

	assert.Equal(product.Name, name)
	assert.Equal(product.Price, price)
	assert.Equal(product.Qtde, qtde)
}

func Test_NewProduct_IDisNotNil(t *testing.T) {
	assert := assert.New(t)

	product, _ := NewProduct(name, price, qtde)

	assert.NotNil(product.ID)
}

func Test_NewProduct_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	product, _ := NewProduct(name, price, qtde)

	assert.Greater(product.CreatedAt, now)
}

func Test_NewProduct_MustValidateName(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewProduct("", price, qtde)

	assert.Equal("name is required", err.Error())
}

func Test_NewProduct_MustValidatePrice(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewProduct(name, 0, qtde)

	assert.Equal("price is required", err.Error())
}

func Test_NewProduct_MustValidateQtde(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewProduct(name, price, 0)

	assert.Equal("qtde is required", err.Error())
}