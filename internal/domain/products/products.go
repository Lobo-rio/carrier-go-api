package products

import (
	"time"

	"github.com/rs/xid"
)

type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64    `json:"price"`
	Qtde      int     `json:"qtde"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func NewProduct(name string, price float64, qtde int) (*Product, error) {
	return &Product{
		ID:        xid.New().String(),
		Name:      name,
		Price:     price,
		Qtde:      qtde,
		CreatedAt: time.Now(),
	}, nil
}