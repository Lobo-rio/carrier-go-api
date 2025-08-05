package clients

import (
	internalerrors "carrierCheck/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)
type AddressClients struct {
	ID 	  string `json:"id" gorm:"size:50"`
	Address string `json:"address" validate:"min=7,max=100,required" gorm:"size:100"`
	Number string `json:"number" validate:"required" gorm:"size:6"`
	Complement string `json:"complement" gorm:"size:20"`
	Neighborhood string `json:"neighborhood" validate:"required" gorm:"size:30"`
	City string `json:"city" validate:"required" gorm:"size:30"`
	State string `json:"state" validate:"required" gorm:"size:2"`
	ClientId string `json:"client_id" gorm:"size:50"`
}
type Client struct {
	ID        string `json:"id" validate:"required" gorm:"size:50"`
	Name      string `json:"name" validate:"min=3,max=60" gorm:"size:60"`
	Email     string `json:"email" validate:"email" gorm:"size:100"`
	Phone     string `json:"phone" validate:"required" gorm:"size:27"`
	Address   []AddressClients `gorm:"foreignKey:ClientId;constraint:OnDelete:CASCADE;" validate:"min=1,dive"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt string `json:"updated_at"`
}

func NewClient(name string, email string, phone string, addressClients []AddressClients) (*Client, error) {
	addressClient := make([]AddressClients, len(addressClients))

	for index, address := range addressClients {
		addressClient[index].Address = address.Address
		addressClient[index].Number = address.Number
		addressClient[index].Complement = address.Complement
		addressClient[index].Neighborhood = address.Neighborhood
		addressClient[index].City = address.City
		addressClient[index].State = address.State
		addressClient[index].ID = xid.New().String()
	}

	client := &Client{
		ID:        xid.New().String(),
		Name:      name,
		Email:    email,
		Phone:     phone,
		Address:  addressClient,
		CreatedAt: time.Now(),
	}
	
	err := internalerrors.ValidateStruct(client)
    if err == nil {
	    return client, nil
    }

	return nil, err
}