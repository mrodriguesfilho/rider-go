package database

import (
	"fmt"
	"rider-go/internal/entity"

	"github.com/google/uuid"
)

type AccountRepositoryInMemory struct {
	Db []entity.Account
}

func NewAccountRepository(db []entity.Account) *AccountRepositoryInMemory {
	return &AccountRepositoryInMemory{
		Db: db,
	}
}

func (a *AccountRepositoryInMemory) Insert(account *entity.Account) error {

	a.Db = append(a.Db, *account)

	return nil
}

func (a *AccountRepositoryInMemory) GetById(id uuid.UUID) (entity.Account, error) {
	for i := 0; i < len(a.Db); i++ {
		if a.Db[i].Id == id {
			return a.Db[i], nil
		}
	}

	return entity.Account{}, fmt.Errorf("no account with id %d was found", id)
}

func (a *AccountRepositoryInMemory) GetByEmail(email string) (entity.Account, error) {
	for i := 0; i < len(a.Db); i++ {

		if a.Db[i].Email == email {
			return a.Db[i], nil
		}
	}

	return entity.Account{}, fmt.Errorf("no account with email %s was found", email)
}
