package carrier

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
    name   = "Test Carrier"
    email  = []string{"test@example.com", "test2@example.com"}
    phone  = "21 97865-8989"
    contact = "Contact Carrier"
	fake = faker.New()
)

func Test_NewCarrier_Create(t *testing.T) {
	assert := assert.New(t)

    carrier, _ := NewCarrier(name, phone, contact, email)

	assert.Equal(carrier.Name, name)
	assert.Equal(carrier.Phone, phone)
	assert.Equal(carrier.Contact, contact)
	assert.Equal(len(carrier.Email), len(email))

}

func Test_NewCarrier_IDisNotNil(t *testing.T) {
	assert := assert.New(t)

	carrier, _ := NewCarrier(name, phone, contact, email)

	assert.NotNil(carrier.ID)
}

func Test_NewCarrier_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	carrier, _ := NewCarrier(name, phone, contact, email)

	assert.Greater(carrier.CreatedAt, now)
}

func Test_NewCarrier_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCarrier(fake.Lorem().Text(2), phone, contact, email)

	assert.Equal("name is required with min 3", err.Error())
}

func Test_NewCarrier_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCarrier(fake.Lorem().Text(65), phone, contact, email)

	assert.Equal("name is required with max 60", err.Error())
}

func Test_NewCarrier_MustValidatePhone(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCarrier(name, "", contact, email)

	assert.Equal("phone is required", err.Error())
}

func Test_NewCarrier_MustValidateContactMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCarrier(name, phone, fake.Lorem().Text(3), email)

	assert.Equal("contact is required with min 5", err.Error())
}


func Test_NewCarrier_MustValidateContactMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCarrier(name, phone, fake.Lorem().Text(61), email)

	assert.Equal("contact is required with max 60", err.Error())
}
func Test_NewCarrier_MustValidateEmailMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCarrier(name, phone, contact, []string{})

	assert.Equal("email is required with min 1", err.Error())
}