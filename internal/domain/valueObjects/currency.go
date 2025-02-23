package valueObjects

type Money struct {
	value    float64
	currency Currency
}

type Currency int

const (
	BRL Currency = iota
	USD
)

func NewMoney(value float64, currency Currency) Money {
	return Money{
		value:    value,
		currency: currency,
	}
}

func (m Money) Equals(other Money) bool {
	return m.value == other.value && m.currency == other.currency
}

func (m Money) GetValue() float64 {
	return m.value
}

func (m Money) GetCurrency() Currency {
	return m.currency
}
