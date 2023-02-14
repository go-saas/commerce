package price

import (
	"github.com/bojanz/currency"
)

// Price database friendly price struct
type Price struct {
	//Amount minor units (e.g. cents). for some database(like sqlite), decimal is stored as float, it may cause problems when calculating
	Amount       int64 `json:"amount"`
	CurrencyCode string
}

// NewPriceFromCurrency convert currency.Amount into database friendly Price
func NewPriceFromCurrency(a currency.Amount) (Price, error) {
	v, err := a.Int64()
	if err != nil {
		return Price{}, err
	}
	return Price{Amount: v, CurrencyCode: a.CurrencyCode()}, nil
}

func NewPrice(n, currencyCode string) (p Price, err error) {
	amount, err := currency.NewAmount(n, currencyCode)
	if err != nil {
		return
	}
	return NewPriceFromCurrency(amount)
}

func NewPriceFromInt64(n int64, currencyCode string) (p Price, err error) {
	amount, err := currency.NewAmountFromInt64(n, currencyCode)
	if err != nil {
		return
	}
	return NewPriceFromCurrency(amount)
}

func NewPriceFromPb(a *PricePb) (Price, error) {
	return NewPriceFromInt64(a.Amount, a.CurrencyCode)
}

func (p Price) ToCurrency() currency.Amount {
	v, err := currency.NewAmountFromInt64(p.Amount, p.CurrencyCode)
	if err != nil {
		panic(err)
	}
	return v
}

func (p Price) ToPricePb() *PricePb {
	a := p.ToCurrency()
	return &PricePb{
		Amount:       p.Amount,
		CurrencyCode: p.CurrencyCode,
		Text:         a.String(),
	}
}
