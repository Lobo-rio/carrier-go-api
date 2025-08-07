package clients

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
    name   = "Test Client"
    email  = "test@example.com"
    phone  = "21 97865-8989"
	fake = faker.New()
)

func Test_NewClient_Create(t *testing.T) {
    assert := assert.New(t)

	client, _ := NewClient(name, email, phone, []AddressClients{})

	assert.Equal(client.Name, name)
	assert.Equal(client.Email, email)
	assert.Equal(client.Phone, phone)
}

func Test_NewClient_IDisNotNil(t *testing.T) {
	assert := assert.New(t)

	client, _ := NewClient(name, email, phone, []AddressClients{})

	assert.NotNil(client.ID)
}

func Test_NewClient_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	client, _ := NewClient(name, email, phone, []AddressClients{})

	assert.Greater(client.CreatedAt, now)
}

func Test_NewClient_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewClient(fake.Lorem().Text(2), phone, email, []AddressClients{})

	assert.Equal("name is required with min 3", err.Error())
}

func Test_NewClient_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewClient(fake.Lorem().Text(61), email, phone, []AddressClients{})

	assert.Equal("name is required with max 60", err.Error())
}

func Test_NewClient_MustValidatePhone(t *testing.T) {
	assert := assert.New(t)

	_, err := NewClient(name, email, "", []AddressClients{})

	assert.Equal("phone is required", err.Error())
}

func Test_NewClient_MustValidateEmail(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewClient(name, "email-invalid", phone, []AddressClients{})

	assert.Equal("email is invalid", err.Error())
}