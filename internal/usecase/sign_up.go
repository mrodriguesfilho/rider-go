package usecase

import (
	"errors"
	"fmt"
	"rider-go/internal/entity"
	"rider-go/internal/infra/database"
)

type SignUpUseCase struct {
	AccountRepository database.AccountRepository
}

type SignUpInput struct {
	Name        string
	Cpf         string
	Email       string
	Password    string
	IsPassenger bool
	IsDriver    bool
}

type SignUpOutput struct {
	Id          int
	Name        string
	Cpf         string
	Email       string
	IsPassenger bool
	IsDriver    bool
}

func NewSignUpUseCase(accountRepository database.AccountRepository) *SignUpUseCase {
	return &SignUpUseCase{
		AccountRepository: accountRepository,
	}
}

func (s *SignUpUseCase) Execute(input SignUpInput) (SignUpOutput, error) {

	accountAlreadyExist, _ := s.AccountRepository.GetByEmail(input.Email)

	if accountAlreadyExist.Id != 0 {
		errorMsg := fmt.Sprintf("%s already signed up on our database", input.Email)
		return SignUpOutput{}, errors.New(errorMsg)
	}

	account := entity.NewAccount(input.Name, input.Cpf, input.Email, input.Password, input.IsPassenger, input.IsDriver)

	s.AccountRepository.Insert(account)

	return SignUpOutput{
		Id:          account.Id,
		Name:        account.Name,
		Cpf:         account.Cpf,
		Email:       account.Email,
		IsPassenger: account.IsPassenger,
		IsDriver:    account.IsDriver,
	}, nil
}
