package contracts

type OrderProduct struct {
	ProductId string
	Quantity  int
	Price     float64
}

type CreateOrder struct {
	ClientId  string
	AddressId string
	Products  []OrderProduct
}