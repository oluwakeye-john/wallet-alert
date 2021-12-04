package currencies

import "github.com/oluwakeye-john/wallet-alert/utils/validators"

type Currency struct {
	Name    string
	Code    string
	IsValid func(string) bool
}

var SupportedCurrencies = []Currency{
	{
		Name:    "Bitcoin",
		Code:    "BTC",
		IsValid: validators.VerifyBTCAddress,
	},
	{
		Name:    "Ethereum",
		Code:    "ETH",
		IsValid: validators.VerifyETHAddress,
	},
	{
		Name:    "Litecoin",
		Code:    "LTC",
		IsValid: validators.VerifyLTCAddress,
	},
	{
		Name:    "DogeCoin",
		Code:    "DOGE",
		IsValid: validators.VerifyDOGEAddress,
	},
	{
		Name:    "Dash",
		Code:    "DASH",
		IsValid: validators.VerifyDASHAddress,
	},
	{
		Name:    "Bcy",
		Code:    "BCY",
		IsValid: validators.VerifyTestAddress,
	},
}
