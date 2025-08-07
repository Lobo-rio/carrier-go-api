package products

import (
	contracts "carrierCheck/internal/contracts/products"
	internalerrors "carrierCheck/internal/internal-errors"
)

type ProductsService interface {
	Create(createProduct contracts.CreateProduct) (string, error)
	GetAll() ([]contracts.ResponseProduct, error)
	GetById(id string) (*contracts.ResponseProduct, error)
	Update(id string, request contracts.UpdateProduct) error
	Delete(id string)  error
}

type ProductsServiceImp struct {
	Repository ProductsRepository
}

func (s *ProductsServiceImp) Create(products contracts.CreateProduct) (string, error) {
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

func (s *ProductsServiceImp) GetAll() ([]contracts.ResponseProduct, error) {
	products, err := s.Repository.GetAll()
	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	responseProducts := make([]contracts.ResponseProduct, len(products))
	for i, product := range products {
		responseProducts[i] = contracts.ResponseProduct{
			ID:        product.ID,
			Name:      product.Name,
			Price:      product.Price,
			Qtde:      product.Qtde,
			CreatedAt: product.CreatedAt.String(),
		}
	}

	return responseProducts, nil
}

func (s *ProductsServiceImp) GetById(id string) (*contracts.ResponseProduct, error) {
	product, err := s.Repository.GetById(id)
	if err != nil {
		return nil, internalerrors.ProcessErrorToReturn(err)
	}
	return &contracts.ResponseProduct{
		ID:        product.ID,
		Name:      product.Name,
		Price:      product.Price,
		Qtde:      product.Qtde,
		CreatedAt: product.CreatedAt.String(),
	}, nil
}

func (s *ProductsServiceImp) Update(id string, request contracts.UpdateProduct) error {
	product, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Qtde = request.Qtde

	err = s.Repository.Update(product)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}

func (s *ProductsServiceImp) Delete(id string)  error {
	product, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	err = s.Repository.Delete(product)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}