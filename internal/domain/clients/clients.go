package clients

import (
	internalerrors "carrierCheck/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Client struct {
	ID        string `json:"id" validate:"required" gorm:"size:50"`
	Name      string `json:"name" validate:"min=3,max=60" gorm:"size:60"`
	Email     string `json:"email" validate:"email" gorm:"size:100"`
	Phone     string `json:"phone" validate:"required" gorm:"size:27"`
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