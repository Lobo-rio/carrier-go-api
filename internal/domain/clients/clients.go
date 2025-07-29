package clients

import (
	internalerrors "carrierCheck/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Client struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"min=3,max=60"`
	Email     string `json:"email" validate:"email"`
	Phone     string `json:"phone" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt string `json:"updated_at"`
}

func NewClient(name string, email string, phone string) (*Client, error) {
	client := &Client{
		ID:        xid.New().String(),
		Name:      name,
		Email:    email,
		Phone:     phone,
		CreatedAt: time.Now(),
	}
	
	err := internalerrors.ValidateStruct(client)
    if err == nil {
	    return client, nil
    }

	return nil, err
}