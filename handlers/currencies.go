package handlers

import (
	"context"

	"github.com/oluwakeye-john/wallet-alert/currencies"
	"github.com/oluwakeye-john/wallet-alert/graph/model"
)

func GetSupportedCurrencies(ctx context.Context) ([]*model.Currency, error) {
	currency_list := []*model.Currency{}

	for _, i := range currencies.SupportedCurrencies {
		currency := &model.Currency{
			Code: model.CurrencyCode(i.Code),
			Name: i.Name,
		}
		currency_list = append(currency_list, currency)
	}

	return currency_list, nil
}
