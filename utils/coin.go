package utils

import (
	"github.com/oluwakeye-john/wallet-alert/currencies"
	"github.com/oluwakeye-john/wallet-alert/customerrors"
)

// This function throws an error if an unsupported coin is passed into it
func GetCoin(code string) (string, error) {
	coin := ""
	for _, i := range currencies.SupportedCurrencies {
		if i.Code == code {
			coin = code
		}
	}

	if coin == "" {
		return "", customerrors.UnsupportedCurrency(coin)
	}

	return coin, nil
}

func GetChain(code string) string {
	if code == "BCY" {
		return "test"
	}

	return "main"
}
