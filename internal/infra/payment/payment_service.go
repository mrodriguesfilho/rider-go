package payment

import "rider-go/internal/domain/valueObjects"

type PaymentService interface {
	Debit(email string, amount valueObjects.Money) bool
	Credit(email string, amount valueObjects.Money) bool
}
