package carrier

import (
	internalerrors "carrierCheck/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type EmailCarrier struct {
	ID 	  string `json:"id" gorm:"size:50"`
	Email string `json:"email" validate:"email,min=7,max=100,required" gorm:"size:100"`
	CarrierId string `json:"carrier_id" gorm:"size:50"`
}

type Carrier struct {
    ID        string         `json:"id" validate:"required" gorm:"size:50"`
    Name      string         `json:"name" validate:"min=3,max=60" gorm:"size:60"`
    Email     []EmailCarrier `gorm:"foreignKey:CarrierId;constraint:OnDelete:CASCADE;" validate:"min=1,dive"`
    Phone     string         `json:"phone" validate:"min=13,max=27" gorm:"size:27"`
    Contact   string         `json:"contact" validate:"min=5,max=60" gorm:"size:60"`
    CreatedAt time.Time      `json:"created_at" validate:"required"`
    UpdatedAt string         `json:"updated_at"`
}

func NewCarrier(name string, phone string, contact string, emails []string) (*Carrier, error) {
	emailCarrier := make([]EmailCarrier, len(emails))

	for index, email := range emails {
		emailCarrier[index].Email = email
		emailCarrier[index].ID = xid.New().String()
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