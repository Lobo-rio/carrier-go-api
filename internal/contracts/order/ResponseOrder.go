package contracts

type ResponseOrder struct {
	ID        string
	ClientId  string
	AddressId string
	CarrierId string
	Status    string
	Products  interface{}
	CreatedAt string
	UpdatedAt string
}