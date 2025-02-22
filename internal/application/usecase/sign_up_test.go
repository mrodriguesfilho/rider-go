package usecase

import (
	"rider-go/internal/domain/entity"
	inmemory "rider-go/internal/infra/database/InMemory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUpUseCase(t *testing.T) {
	t.Run("It should create a valid account", func(t *testing.T) {
		signUpInput := SignUpInput{
			Name:        "John Doe",
			Email:       "johndoe@gmail.com",
			Password:    "123123",
			IsPassenger: true,
			IsDriver:    false,
		}

		accountRepository := inmemory.NewAccountRepository(make([]entity.Account, 0))
		signUpUseCase := NewSignUpUseCase(accountRepository)
		signUpOutput, err := signUpUseCase.Execute(signUpInput)

		assert.Nil(t, err)
		assert.NotNil(t, signUpOutput)
		assert.NotEqual(t, 0, signUpOutput.Id)
	})

	t.Run("It shouldn't create an account with e-mail already used", func(t *testing.T) {
		signUpInput := SignUpInput{
			Name:        "John Doe",
			Email:       "johndoe@gmail.com",
			Password:    "123123",
			IsPassenger: true,
			IsDriver:    false,
		}

		accountRepository := inmemory.NewAccountRepository(make([]entity.Account, 0))
		signUpUseCase := NewSignUpUseCase(accountRepository)
		signUpOutput, err := signUpUseCase.Execute(signUpInput)

		assert.Nil(t, err)
		assert.NotNil(t, signUpOutput)
		assert.NotEqual(t, 0, signUpOutput.Id)

		signUpOutput, err = signUpUseCase.Execute(signUpInput)

		assert.NotNil(t, err)
		assert.Equal(t, "johndoe@gmail.com already signed up on our database", err.Error())

	})
}
