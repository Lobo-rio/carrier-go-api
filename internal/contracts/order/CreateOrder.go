package contracts

type Product struct {
	ProductId string
	Qtde      int
	Price     float64
}

type CreateOrder struct {
	ClientId  string
	AddressId string
	Products  []Product
}