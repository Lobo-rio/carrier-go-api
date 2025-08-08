package products

import (
	internalerrors "carrierCheck/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Product struct {
	ID        string    `json:"id" validate:"required" gorm:"size:50;primaryKey"`
	Name      string    `json:"name" validate:"min=5,max=80" gorm:"size:80"`
	Price     float64    `json:"price" validate:"required,gt=0"`
	Qtde      int     `json:"qtde" validate:"required,gt=0"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt string    `json:"updated_at"`
}

func NewProduct(name string, price float64, qtde int) (*Product, error) {
	product := &Product{
		ID:        xid.New().String(),
		Name:      name,
		Price:     price,
		Qtde:      qtde,
		CreatedAt: time.Now(),
	}
	
	err := internalerrors.ValidateStruct(product)
    if err == nil {
	    return product, nil
    }

	return nil, err
}