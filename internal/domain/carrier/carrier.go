package carrier

import (
	"time"

	"github.com/rs/xid"
)

type EmailCarrier struct {
	Email string `json:"email"`
}

type Carrier struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     []EmailCarrier
	Phone     string `json:"phone"`
	Contact   string `json:"contact"`
	CreatedAt time.Time
	UpdatedAt string `json:"updated_at"`
}

func NewCarrier(name string, phone string, contact string, emails []string) (*Carrier, error) {
	emailCarrier := make([]EmailCarrier, len(emails))

	for index, email := range emails {
		emailCarrier[index].Email = email
	}

	return &Carrier{
		ID:        xid.New().String(),
		Name:      name,
		Email:    emailCarrier,
		Phone:     phone,
		Contact:   contact,
		CreatedAt: time.Now(),
	}, nil
}