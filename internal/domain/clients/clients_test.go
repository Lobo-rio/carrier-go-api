package clients

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
    name   = "Test Client"
    email  = "test@example.com"
    phone  = "21 97865-8989"
)

func Test_NewClient_Create(t *testing.T) {
    assert := assert.New(t)

	client, _ := NewClient(name, phone, email)

	assert.Equal(client.Name, name)
	assert.Equal(client.Phone, phone)
	assert.Equal(len(client.Email), len(email))

}

func Test_NewClient_IDisNotNil(t *testing.T) {
	assert := assert.New(t)

	client, _ := NewClient(name, phone, email)

	assert.NotNil(client.ID)
}

func Test_NewClient_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	client, _ := NewClient(name, phone, email)

	assert.Greater(client.CreatedAt, now)
}

func Test_NewClient_MustValidateName(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewClient("", phone, email)

	assert.Equal("name is required", err.Error())
}

func Test_NewClient_MustValidatePhone(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewClient(name, "", email)

	assert.Equal("phone is required", err.Error())
}

func Test_NewClient_MustValidateEmail(t *testing.T) {
	assert := assert.New(t)
	
	_, err := NewClient(name, phone, "")

	assert.Equal("email is required", err.Error())
}