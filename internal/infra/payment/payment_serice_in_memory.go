package payment

import (
	"rider-go/internal/domain/valueObjects"
	"time"
)

type PaymentInMemoryDTO struct {
	Amount valueObjects.Money
	Type   PaymentTypeInMemory
	Date   time.Time
}

type PaymentTypeInMemory int

const (
	None PaymentTypeInMemory = iota
	Credit
	Debit
)

type PaymentServiceInMemoryAdapter struct {
	Payments map[string]PaymentInMemoryDTO
}

func NewPaymentServiceInMemory() *PaymentServiceInMemoryAdapter {
	return &PaymentServiceInMemoryAdapter{
		Payments: make(map[string]PaymentInMemoryDTO),
	}
}

func (p *PaymentServiceInMemoryAdapter) Debit(email string, amount valueObjects.Money) bool {

	date := time.Now()

	paymentInMemoryDTO := PaymentInMemoryDTO{
		Amount: amount,
		Date:   date,
		Type:   Debit,
	}

	p.Payments[email] = paymentInMemoryDTO

	return true
}

func (p *PaymentServiceInMemoryAdapter) Credit(email string, amount valueObjects.Money) bool {
	date := time.Now()

	paymentInMemoryDTO := PaymentInMemoryDTO{
		Amount: amount,
		Date:   date,
		Type:   Credit,
	}

	p.Payments[email] = paymentInMemoryDTO

	return true
}
