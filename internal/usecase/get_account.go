package usecase

import "rider-go/internal/infra/database"

type GetAccountUsecase struct {
	accountRepository database.AccountRepository
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

func NewGetAccountUseCase(accountRepository database.AccountRepository) *GetAccountUsecase {
	return &GetAccountUsecase{
		accountRepository: accountRepository,
	}
}

func (g *GetAccountUsecase) Execute(getAccountInput GetAccountInput) (GetAccountOutput, error) {

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
