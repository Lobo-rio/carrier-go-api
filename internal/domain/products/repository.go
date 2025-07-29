package products

type ProductsRepository interface {
	Save(product *Product) error
}