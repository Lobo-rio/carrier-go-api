package orders

import (
	"carrierCheck/internal/domain/clients"
	"carrierCheck/internal/domain/products"
	internalerrors "carrierCheck/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Product struct {
	ProductId string  `json:"product_id" gorm:"size:50"`
	Product   products.Product `json:"product" gorm:"foreignKey:ProductId;references:ID"`
	OrderId   string  `json:"order_id" gorm:"size:50"`
	Quantity  int     `json:"quantity"`
	Price	  float64  `json:"price"` 
}

type Order struct {
	ID              string    `json:"id" gorm:"size:50"`
	ClientId       string    `json:"clients_id" gorm:"size:50"`
	Client          clients.Client `json:"client" gorm:"foreignKey:ClientsId;references:ID"`
	AddressId       string    `json:"address_id" gorm:"size:50"`
	Address         clients.AddressClients  `json:"address" gorm:"foreignKey:AddressId;references:ID"`
	Products        []Product `json:"products" gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE;"`
	Status          string    `json:"status" gorm:"size:20"`
	CreatedAt       time.Time `json:"created_at" validate:"required"`
	UpdatedAt       string   `json:"updated_at"` 
}

func NewOrder(clientId string, AddressId string, products []Product) (*Order, error) {
	product := make([]Product, len(products))

	for index, prd := range products {
		product[index].ProductId = prd.ProductId
		product[index].Quantity = prd.Quantity
		product[index].Price = prd.Price
	}

	order := &Order{
		ID:        xid.New().String(),
		ClientId:  clientId,
 		AddressId: AddressId,
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