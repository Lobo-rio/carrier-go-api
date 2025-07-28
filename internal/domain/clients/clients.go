package clients

import (
	"time"

	"github.com/rs/xid"
)

type Client struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewClient(name string, email string, phone string) (*Client, error) {
	return &Client{
		ID:        xid.New().String(),
		Name:      name,
		Email:    email,
		Phone:     phone,
		CreatedAt: time.Now(),
	}, nil
}