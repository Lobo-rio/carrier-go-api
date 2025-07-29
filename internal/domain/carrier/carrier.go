package carrier

import (
	internalerrors "carrierCheck/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type EmailCarrier struct {
	Email string `json:"email" validate:"email,required"`
}

type Carrier struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"min=3,max=60"`
	Email     []EmailCarrier `validate:"min=1,dive"`
	Phone     string `json:"phone" validate:"required"`
	Contact   string `json:"contact" validate:"min=5,max=60"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt string `json:"updated_at"`
}

func NewCarrier(name string, phone string, contact string, emails []string) (*Carrier, error) {
	emailCarrier := make([]EmailCarrier, len(emails))

	for index, email := range emails {
		emailCarrier[index].Email = email
	}

	carrier := &Carrier{
		ID:        xid.New().String(),
		Name:      name,
		Email:    emailCarrier,
		Phone:     phone,
		Contact:   contact,
		CreatedAt: time.Now(),
	}

	err := internalerrors.ValidateStruct(carrier)
    if err == nil {
	    return carrier, nil
    }

	return nil, err
}