package contracts

type AddressClient struct {
	Address      string
	Number       string
	Complement   string
	Neighborhood string
	City         string
	State        string
}

type CreateClient struct {
	Name    string
	Address []AddressClient
	Email   string
	Phone   string
}