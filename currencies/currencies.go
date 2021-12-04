package currencies

import (
	"strings"

	"github.com/oluwakeye-john/wallet-alert/customerrors"
	"github.com/oluwakeye-john/wallet-alert/utils/validators"
)

type Currency struct {
	Name    string
	Code    string
	IsValid func(string) bool
	Chain   string
}

var TestCurrency = Currency{
	Name:    "Bcy",
	Code:    "BCY",
	IsValid: validators.VerifyTestAddress,
	Chain:   "test",
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
		Name:    "Bitcoin",
		Code:    "BTC",
		IsValid: validators.VerifyBTCAddress,
		Chain:   "main",
	},
	{
		Name:    "Ethereum",
		Code:    "ETH",
		IsValid: validators.VerifyETHAddress,
		Chain:   "main",
	},
	{
		Name:    "Litecoin",
		Code:    "LTC",
		IsValid: validators.VerifyLTCAddress,
		Chain:   "main",
	},
	{
		Name:    "DogeCoin",
		Code:    "DOGE",
		IsValid: validators.VerifyDOGEAddress,
		Chain:   "main",
	},
	{
		Name:    "Dash",
		Code:    "DASH",
		IsValid: validators.VerifyDASHAddress,
		Chain:   "main",
	},
	TestCurrency,
}
