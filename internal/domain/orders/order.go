package orders

import (
	internalerrors "carrierCheck/internal/internal-errors"
	"fmt"
	"time"

	"github.com/rs/xid"
)

type OrderProduct struct {
	ID    string `json:"id" gorm:"size:50;primaryKey"`
	ProductId string  `json:"product_id" gorm:"size:50" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	Price	  float64  `json:"price" validate:"required"` 
	OrderId   string  `json:"order_id" gorm:"size:50"`
}

type Order struct {
	ID              string    `json:"id" gorm:"size:50;primaryKey"`
	ClientId       string    `json:"client_id" gorm:"size:50" validate:"required"`
	AddressId       string    `json:"address_id" gorm:"size:50" validate:"required"`
	Status          string    `json:"status" gorm:"size:20"`
	Products        []OrderProduct `json:"products" gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE;"  validate:"min=1,dive"`
	CreatedAt       time.Time `json:"created_at" validate:"required"`
	UpdatedAt       string   `json:"updated_at"` 
}

func NewOrder(clientId string, addressId string, products []OrderProduct) (*Order, error) {
	product := make([]OrderProduct, len(products))
	for index, prd := range products {
		product[index] = OrderProduct{
			ID:        xid.New().String(),
			ProductId: prd.ProductId,
			Quantity:  prd.Quantity,
			Price:    prd.Price,
		}
	}

	fmt.Println("Validation product:", product)

	order := &Order{
		ID:        xid.New().String(),
		ClientId:  clientId,
		AddressId: addressId,
		Products:  product,
		Status:	   "pending",
		CreatedAt: time.Now(),
	}

	err := internalerrors.ValidateStruct(order)
	if err == nil {
		return order, nil
	}

	return nil, err
}