package price

import "github.com/shopspring/decimal"

type Price struct {
	Amount       decimal.Decimal `json:"amount" sql:"type:decimal(20,8);"`
	CurrencyCode string
}
