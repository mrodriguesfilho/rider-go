package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	t.Run("It should create a valid account", func(t *testing.T) {
		account := NewAccount("John Doe", "johndoe@gmail.com", "123123", true, false)

		assert.Empty(t, account.Validation.Errors)
		assert.Equal(t, "John Doe", account.Name)
		assert.Equal(t, "johndoe@gmail.com", account.Email)
		assert.Equal(t, "123123", account.Password)
		assert.True(t, account.IsPassenger)
		assert.False(t, account.IsDriver)
	})

	t.Run("It shouldn't create an account without name", func(t *testing.T) {
		account := NewAccount("", "johndoe@gmail.com", "123123", true, false)

		assert.NotEmpty(t, account.Validation.Errors)
		assert.Len(t, account.Validation.Errors, 1)

		actualErrorMessage := account.Validation.Errors[0].Error()
		assert.Equal(t, "name property is required for Account type", actualErrorMessage)
	})

	t.Run("It shouldn't create an account without email", func(t *testing.T) {
		account := NewAccount("John Doe", "", "123123", true, false)

		assert.NotEmpty(t, account.Validation.Errors)
		assert.Len(t, account.Validation.Errors, 1)

		actualErrorMessage := account.Validation.Errors[0].Error()
		assert.Equal(t, "email property is required for Account type", actualErrorMessage)
	})

	t.Run("It shouldn't create an account without password", func(t *testing.T) {
		account := NewAccount("John Doe", "johndoe@gmail.com", "", true, false)

		assert.NotEmpty(t, account.Validation.Errors)
		assert.Len(t, account.Validation.Errors, 1)

		actualErrorMessage := account.Validation.Errors[0].Error()
		assert.Equal(t, "password property is required for Account type", actualErrorMessage)
	})

	t.Run("It shouldn't create an account without empty data", func(t *testing.T) {
		account := NewAccount("", "", "", true, false)

		assert.NotEmpty(t, account.Validation.Errors)
		assert.Len(t, account.Validation.Errors, 3)

		assert.Equal(t, "name property is required for Account type", account.Validation.Errors[0].Error())
		assert.Equal(t, "email property is required for Account type", account.Validation.Errors[1].Error())
		assert.Equal(t, "password property is required for Account type", account.Validation.Errors[2].Error())
	})
}
