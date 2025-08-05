package products

type ProductsRepository interface {
	Save(product *Product) error
	GetAll() ([]Product, error)
	GetById(id string) (*Product, error)
	Update(product *Product) error
	Delete(product *Product) error
}