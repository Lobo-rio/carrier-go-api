package carrier

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
    name   = "Test Carrier"
    email  = []string{"test@example.com", "test2@example.com"}
    phone  = "21 97865-8989"
    contact = "Contact Carrier"
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

func Test_NewCarrier_MustValidateName(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewCarrier("", phone, contact, email)

	assert.Equal("name is required", err.Error())
}

func Test_NewCarrier_MustValidatePhone(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewCarrier(name, "", contact, email)

	assert.Equal("phone is required", err.Error())
}

func Test_NewCarrier_MustValidateContact(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewCarrier(name, phone, "", email)

	assert.Equal("contact is required", err.Error())
}

func Test_NewCarrier_MustValidateEmail(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewCarrier(name, phone, contact, []string{})

	assert.Equal("email is required", err.Error())
}