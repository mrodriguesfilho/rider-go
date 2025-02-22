package repository

import (
	"rider-go/internal/domain/entity"

	"github.com/google/uuid"
)

type AccountRepository interface {
	Insert(account *entity.Account) error
	GetByEmail(email string) (entity.Account, error)
	GetById(id uuid.UUID) (entity.Account, error)
}
