package database

import (
	"fmt"
	"rider-go/internal/entity"

	"github.com/google/uuid"
)

type AccountRepositoryInMemory struct {
	db []entity.Account
}

func NewAccountRepository(db []entity.Account) *AccountRepositoryInMemory {
	return &AccountRepositoryInMemory{
		db: db,
	}
}

func (a *AccountRepositoryInMemory) Insert(account *entity.Account) error {

	a.db = append(a.db, *account)

	return nil
}

func (a *AccountRepositoryInMemory) GetById(id uuid.UUID) (entity.Account, error) {
	for i := 0; i < len(a.db); i++ {
		if a.db[i].Id == id {
			return a.db[i], nil
		}
	}

	return entity.Account{}, fmt.Errorf("no account with id %d was found", id)
}

func (a *AccountRepositoryInMemory) GetByEmail(email string) (entity.Account, error) {
	for i := 0; i < len(a.db); i++ {

		if a.db[i].Email == email {
			return a.db[i], nil
		}
	}

	return entity.Account{}, fmt.Errorf("no account with email %s was found", email)
}
