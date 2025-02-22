package usecase

import (
	"errors"
	"fmt"
	"rider-go/internal/domain/entity"
	"rider-go/internal/infra/database/repository"

	"github.com/google/uuid"
)

type SignUpUseCase struct {
	accountRepository repository.AccountRepository
}

type SignUpInput struct {
	Name        string
	Email       string
	Password    string
	IsPassenger bool
	IsDriver    bool
}

type SignUpOutput struct {
	Id          string
	Name        string
	Email       string
	IsPassenger bool
	IsDriver    bool
}

func NewSignUpUseCase(accountRepository repository.AccountRepository) *SignUpUseCase {
	return &SignUpUseCase{
		accountRepository: accountRepository,
	}
}

func (s *SignUpUseCase) Execute(input SignUpInput) (SignUpOutput, error) {

	accountAlreadyExist, _ := s.accountRepository.GetByEmail(input.Email)

	if accountAlreadyExist.Id != uuid.Nil {
		errorMsg := fmt.Sprintf("%s already signed up on our database", input.Email)
		return SignUpOutput{}, errors.New(errorMsg)
	}

	account := entity.NewAccount(input.Name, input.Email, input.Password, input.IsPassenger, input.IsDriver)

	s.accountRepository.Insert(account)

	return SignUpOutput{
		Id:          account.Id.String(),
		Name:        account.Name,
		Email:       account.Email,
		IsPassenger: account.IsPassenger,
		IsDriver:    account.IsDriver,
	}, nil
}
