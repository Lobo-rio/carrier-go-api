package products

import (
	contracts "carrierCheck/internal/contracts/products"
	internalerrors "carrierCheck/internal/internal-errors"
)

type ProductsService struct {
	Repository ProductsRepository
}

func (s *ProductsService) Create(products contracts.CreateProduct) (string, error) {
	product, err := NewProduct(products.Name, products.Price, products.Qtde)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(product)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return product.ID, nil
}