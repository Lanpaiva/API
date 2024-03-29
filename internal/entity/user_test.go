package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Alan Doe", "alan@teste.com", "Senha123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Alan Doe", user.Name)
	assert.Equal(t, "alan@teste.com", user.Email)

}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "Senha123456@")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("Senha123456@"))
	assert.False(t, user.ValidatePassword("Senha1234567"))
	assert.NotEqual(t, "Senha123456@", user.Password)
}
