package usecase

import (
	"rider-go/internal/domain/entity"
	inmemory "rider-go/internal/infra/database/InMemory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountUseCase(t *testing.T) {
	t.Run("It should retrieve a valid account", func(t *testing.T) {
		signUpInput := SignUpInput{
			Name:        "John Doe",
			Email:       "johndoe@gmail.com",
			Password:    "123123",
			IsPassenger: true,
			IsDriver:    false,
		}

		accountRepository := inmemory.NewAccountRepository(make([]entity.Account, 0))
		signUpUseCase := NewSignUpUseCase(accountRepository)
		signUpOutput, signUpErr := signUpUseCase.Execute(signUpInput)

		getAccountUseCase := NewGetAccountUseCase(accountRepository)

		getAccountInput := GetAccountInput{
			Email: signUpInput.Email,
		}

		getAccountOutput, getAccountErr := getAccountUseCase.Execute(getAccountInput)

		assert.Nil(t, signUpErr)
		assert.Nil(t, getAccountErr)
		assert.NotNil(t, signUpOutput)
		assert.NotEqual(t, 0, signUpOutput.Id)
		assert.Equal(t, getAccountOutput.Id, signUpOutput.Id)
		assert.Equal(t, getAccountOutput.Email, signUpOutput.Email)
	})

	t.Run("It shouldn't retrieve an account that doesn't exists", func(t *testing.T) {

		accountRepository := inmemory.NewAccountRepository(make([]entity.Account, 0))

		getAccountUseCase := NewGetAccountUseCase(accountRepository)

		getAccountInput := GetAccountInput{
			Email: "johndoe@gmail.com",
		}

		_, getAccountErr := getAccountUseCase.Execute(getAccountInput)

		assert.NotNil(t, getAccountErr)
		assert.Equal(t, getAccountErr.Error(), "no account with email johndoe@gmail.com was found")
	})
}
