package handlers

import (
	"context"

	"github.com/oluwakeye-john/wallet-alert/currencies"
	"github.com/oluwakeye-john/wallet-alert/customerrors"
	"github.com/oluwakeye-john/wallet-alert/graph/model"
	"github.com/oluwakeye-john/wallet-alert/utils/validators"
)

func CreateSubscription(ctx context.Context, input model.SubscriptionInput) (*model.SubscriptionStatus, error) {
	subscription_status := &model.SubscriptionStatus{}

	is_address_valid := false
	is_currency_supported := false
	is_email_valid := validators.IsEmailValid(input.Email)
	is_name_valid := validators.IsNameValid(input.Name)

	if !is_name_valid {
		return nil, customerrors.InvalidName()
	}

	if !is_email_valid {
		return nil, customerrors.InvalidEmail()
	}

	for _, i := range currencies.SupportedCurrencies {
		if input.CurrencyCode == model.CurrencyCode(i.Code) {
			is_currency_supported = true
			is_address_valid = i.IsValid(input.Address)
			break
		}
	}

	if !is_currency_supported {
		return nil, customerrors.UnsupportedCurrency(string(input.CurrencyCode))
	} else if !is_address_valid {
		return nil, customerrors.InvalidAddress()
	}

	return subscription_status, nil
}
