package currencies

import (
	"fmt"
	"strings"

	"github.com/oluwakeye-john/wallet-alert/customerrors"
	"github.com/oluwakeye-john/wallet-alert/utils/validators"
)

type Currency struct {
	Name            string
	Code            string
	IsAddressValid  func(string) bool
	Chain           string
	ExplorerTxUrl   func(string) string
	ExplorerAddrUrl func(string) string
}

var TestCurrency = Currency{
	Name:           "Bcy",
	Code:           "BCY",
	IsAddressValid: validators.VerifyTestAddress,
	Chain:          "test",
	ExplorerTxUrl: func(s string) string {
		return fmt.Sprintf("https://live.blockcypher.com/bcy/tx/%s", s)
	},
	ExplorerAddrUrl: func(s string) string {
		return fmt.Sprintf("https://live.blockcypher.com/bcy/address/%s", s)
	},
}

func (c *Currency) CodeInLowerCase() string {
	return strings.ToLower(c.Code)
}

func GetCurrencyFromCode(currency_code string) (Currency, error) {
	currency := Currency{}
	for _, i := range SupportedCurrencies {
		if i.Code == currency_code {
			currency = i
		}
	}

	if currency.Code == "" {
		return currency, customerrors.UnsupportedCurrency(currency_code)
	}

	return currency, nil
}

var SupportedCurrencies = []Currency{
	{
		Name:           "Bitcoin",
		Code:           "BTC",
		IsAddressValid: validators.VerifyBTCAddress,
		Chain:          "main",
		ExplorerTxUrl: func(s string) string {
			return ""
		},
		ExplorerAddrUrl: func(s string) string {
			return ""
		},
	},
	{
		Name:           "Ethereum",
		Code:           "ETH",
		IsAddressValid: validators.VerifyETHAddress,
		Chain:          "main",
		ExplorerTxUrl: func(s string) string {
			return ""
		},
		ExplorerAddrUrl: func(s string) string {
			return ""
		},
	},
	{
		Name:           "Litecoin",
		Code:           "LTC",
		IsAddressValid: validators.VerifyLTCAddress,
		Chain:          "main",
		ExplorerTxUrl: func(s string) string {
			return ""
		},
		ExplorerAddrUrl: func(s string) string {
			return ""
		},
	},
	{
		Name:           "DogeCoin",
		Code:           "DOGE",
		IsAddressValid: validators.VerifyDOGEAddress,
		Chain:          "main",
		ExplorerTxUrl: func(s string) string {
			return ""
		},
		ExplorerAddrUrl: func(s string) string {
			return ""
		},
	},
	{
		Name:           "Dash",
		Code:           "DASH",
		IsAddressValid: validators.VerifyDASHAddress,
		Chain:          "main",
		ExplorerTxUrl: func(s string) string {
			return ""
		},
		ExplorerAddrUrl: func(s string) string {
			return ""
		},
	},
	TestCurrency,
}
