package usecase

import (
	"rider-go/internal/infra/database/repository"
)

type GetAccount struct {
	accountRepository repository.AccountRepository
}

type GetAccountInput struct {
	Email string
}

type GetAccountOutput struct {
	Id          string
	Name        string
	Email       string
	IsPassenger bool
	IsDriver    bool
}

func NewGetAccountUseCase(accountRepository repository.AccountRepository) *GetAccount {
	return &GetAccount{
		accountRepository: accountRepository,
	}
}

func (g *GetAccount) Execute(getAccountInput GetAccountInput) (GetAccountOutput, error) {

	account, err := g.accountRepository.GetByEmail(getAccountInput.Email)

	if err != nil {
		return GetAccountOutput{}, err
	}

	return GetAccountOutput{
		Id:          account.Id.String(),
		Name:        account.Name,
		Email:       account.Email,
		IsPassenger: account.IsPassenger,
		IsDriver:    account.IsDriver,
	}, nil

}
