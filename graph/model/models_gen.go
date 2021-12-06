// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Address struct {
	Address      string       `json:"address"`
	PublicKey    string       `json:"public_key"`
	PrivateKey   string       `json:"private_key"`
	CurrencyCode CurrencyCode `json:"currency_code"`
	ExplorerURL  string       `json:"explorer_url"`
}

type CancelSubscriptionInput struct {
	Email string `json:"email"`
}

type CreateSubscriptionInput struct {
	Address      string       `json:"address"`
	Email        string       `json:"email"`
	CurrencyCode CurrencyCode `json:"currency_code"`
}

type Currency struct {
	Code CurrencyCode `json:"code"`
	Name string       `json:"name"`
}

type DeleteHookInput struct {
	HookID       string       `json:"hook_id"`
	CurrencyCode CurrencyCode `json:"currency_code"`
}

type FundTestAddressInput struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

type GetStatusInput struct {
	Email string `json:"email"`
}

type SubscriptionStatus struct {
	IsSubscribed bool   `json:"is_subscribed"`
	Address      string `json:"address"`
}

type Transaction struct {
	Txhash      string  `json:"txhash"`
	Amount      float64 `json:"amount"`
	ExplorerURL string  `json:"explorer_url"`
}

type CurrencyCode string

const (
	CurrencyCodeBtc  CurrencyCode = "BTC"
	CurrencyCodeEth  CurrencyCode = "ETH"
	CurrencyCodeLtc  CurrencyCode = "LTC"
	CurrencyCodeDoge CurrencyCode = "DOGE"
	CurrencyCodeDash CurrencyCode = "DASH"
	CurrencyCodeBcy  CurrencyCode = "BCY"
)

var AllCurrencyCode = []CurrencyCode{
	CurrencyCodeBtc,
	CurrencyCodeEth,
	CurrencyCodeLtc,
	CurrencyCodeDoge,
	CurrencyCodeDash,
	CurrencyCodeBcy,
}

func (e CurrencyCode) IsValid() bool {
	switch e {
	case CurrencyCodeBtc, CurrencyCodeEth, CurrencyCodeLtc, CurrencyCodeDoge, CurrencyCodeDash, CurrencyCodeBcy:
		return true
	}
	return false
}

func (e CurrencyCode) String() string {
	return string(e)
}

func (e *CurrencyCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CurrencyCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CurrencyCode", str)
	}
	return nil
}

func (e CurrencyCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
