// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Currency struct {
	Code CurrencyCode `json:"code"`
	Name string       `json:"name"`
}

type SubscriptionInput struct {
	Address      string       `json:"address"`
	Email        string       `json:"email"`
	Name         string       `json:"name"`
	CurrencyCode CurrencyCode `json:"currency_code"`
}

type SubscriptionStatus struct {
	IsSubscribed bool `json:"is_subscribed"`
}

type CurrencyCode string

const (
	CurrencyCodeBtc  CurrencyCode = "BTC"
	CurrencyCodeEth  CurrencyCode = "ETH"
	CurrencyCodeLtc  CurrencyCode = "LTC"
	CurrencyCodeDoge CurrencyCode = "DOGE"
	CurrencyCodeDash CurrencyCode = "DASH"
)

var AllCurrencyCode = []CurrencyCode{
	CurrencyCodeBtc,
	CurrencyCodeEth,
	CurrencyCodeLtc,
	CurrencyCodeDoge,
	CurrencyCodeDash,
}

func (e CurrencyCode) IsValid() bool {
	switch e {
	case CurrencyCodeBtc, CurrencyCodeEth, CurrencyCodeLtc, CurrencyCodeDoge, CurrencyCodeDash:
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
