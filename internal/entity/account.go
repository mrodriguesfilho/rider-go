package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Account struct {
	Id          uuid.UUID
	Name        string
	Cpf         string
	Email       string
	Password    string
	IsPassenger bool
	IsDriver    bool
	Validation  Validation
}

func NewAccount(name, cpf, email, password string, isPassenger, isDriver bool) *Account {

	account := Account{
		Id:          uuid.New(),
		Name:        name,
		Cpf:         cpf,
		Email:       email,
		Password:    password,
		IsPassenger: isPassenger,
		IsDriver:    isDriver,
	}

	account.Validate()

	return &account
}

func (a *Account) Validate() {

	var validation Validation

	if a.Name == "" {
		validation.Errors = append(validation.Errors, errors.New("name property is required for Account type"))
	}

	if a.Cpf == "" {
		validation.Errors = append(validation.Errors, errors.New("cpf property is required for Account type"))
	}

	if a.Email == "" {
		validation.Errors = append(validation.Errors, errors.New("email property is required for Account type"))
	}

	if a.Password == "" {
		validation.Errors = append(validation.Errors, errors.New("password property is required for Account type"))
	}

	a.Validation = validation
}
