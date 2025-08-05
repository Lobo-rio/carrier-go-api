package contracts

type ResponseOrder struct {
	ID        string
	ClientId  string
	AddressId string
	Products  interface{}
	CreatedAt string
}
