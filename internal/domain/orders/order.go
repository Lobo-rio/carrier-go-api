package orders

import (
	internalerrors "carrierCheck/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type OrdersProducts struct {
	ID        string `json:"id" gorm:"size:50;primaryKey"`
	ProductId string  `json:"product_id" gorm:"size:50" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	Price	  float64  `json:"price" validate:"required"` 
}

type Order struct {
	ID              string    `json:"id" gorm:"size:50;primaryKey"`
	ClientId        string    `json:"client_id" gorm:"size:50" validate:"required"`
	AddressId       string    `json:"address_id" gorm:"size:50" validate:"required"`
	CarrierId	    string     `json:"carrier_id" gorm:"size:50"`
	Status          string    `json:"status" gorm:"size:40"`
	OrderProduct    []OrdersProducts `json:"products" gorm:"many2many:order_orderproduct;"  validate:"min=1,dive"`
	CreatedAt       time.Time `json:"created_at" validate:"required"`
	UpdatedAt       time.Time `json:"updated_at"` 
}

func NewOrder(clientId string, addressId string, products []OrdersProducts) (*Order, error) {
	product := make([]OrdersProducts, len(products))
	for index, prd := range products {
		product[index] = OrdersProducts{
			ID:        xid.New().String(),
			ProductId: prd.ProductId,
			Quantity:  prd.Quantity,
			Price:    prd.Price,
		}
	}

	order := &Order{
		ID:        xid.New().String(),
		ClientId:  clientId,
		AddressId: addressId,
		OrderProduct:  product,
		Status:	   "pending",
		CreatedAt: time.Now(),
	}

	err := internalerrors.ValidateStruct(order)
	if err == nil {
		return order, nil
	}

	return order, nil
}