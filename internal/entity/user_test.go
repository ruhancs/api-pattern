package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user,err := NewUser("test","t@email", "test")

	assert.Nil(t,err)
	assert.NotNil(t,user)
	assert.NotEmpty(t,user.ID)
	assert.NotEmpty(t,user.Password)
	assert.Equal(t,user.Name, "test")
	assert.Equal(t,user.Email, "t@email")
}

func TestUserValidatePassword(t *testing.T) {
	user,err := NewUser("test","t@email", "test")

	assert.Nil(t,err)
	assert.NotNil(t,user)
	assert.True(t,user.ValidatePassword("test"))
	assert.False(t,user.ValidatePassword("12345667"))
	assert.NotEqual(t, user.Password, "test")
}