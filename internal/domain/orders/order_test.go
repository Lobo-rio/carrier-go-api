package orders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder_Success(t *testing.T) {
	products := []OrdersProducts{
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
	}

	order, err := NewOrder("client1", "address1", products)

	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, "client1", order.ClientId)
	assert.Equal(t, "address1", order.AddressId)
	assert.Equal(t, "pending", order.Status)
	assert.Len(t, order.OrderProduct, 2)
	assert.Equal(t, "product1", order.OrderProduct[0].ProductId)
	assert.Equal(t, 2, order.OrderProduct[0].Quantity)
	assert.Equal(t, 100.0, order.OrderProduct[0].Price)
	assert.Equal(t, "product2", order.OrderProduct[1].ProductId)
	assert.Equal(t, 1, order.OrderProduct[1].Quantity)
	assert.Equal(t, 50.0, order.OrderProduct[1].Price)
}

func TestNewOrder_ValidationError(t *testing.T) {
	products := []OrdersProducts{
		{
			ProductId: "",
			Quantity:  0,
			Price:     0.0,
		},
	}

	order, err := NewOrder("", "", products)

	assert.Error(t, err)
	assert.Nil(t, order)
}
