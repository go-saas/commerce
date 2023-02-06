package price

import "github.com/bojanz/currency"

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

func (p Price) ToCurrency() currency.Amount {
	v, _ := currency.NewAmountFromInt64(p.Amount, p.CurrencyCode)
	return v
}
