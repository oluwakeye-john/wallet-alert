package handlers

import (
	"context"
	"errors"

	"github.com/oluwakeye-john/wallet-alert/blockcypher"
	"github.com/oluwakeye-john/wallet-alert/currencies"
	"github.com/oluwakeye-john/wallet-alert/customerrors"
	"github.com/oluwakeye-john/wallet-alert/database"
	"github.com/oluwakeye-john/wallet-alert/graph/model"
	"github.com/oluwakeye-john/wallet-alert/models"
	"github.com/oluwakeye-john/wallet-alert/utils/validators"
	"gorm.io/gorm"
)

func validateInput(input *model.CreateSubscriptionInput) error {
	is_address_valid := false
	is_currency_supported := false
	is_email_valid := validators.IsEmailValid(input.Email)

	if !is_email_valid {
		return customerrors.InvalidEmail()
	}

	for _, i := range currencies.SupportedCurrencies {
		if input.CurrencyCode == model.CurrencyCode(i.Code) {
			is_currency_supported = true
			is_address_valid = i.IsValid(input.Address)
			break
		}
	}

	if !is_currency_supported {
		return customerrors.UnsupportedCurrency(string(input.CurrencyCode))
	} else if !is_address_valid {
		return customerrors.InvalidAddress()
	}

	return nil
}

func CreateSubscription(ctx context.Context, input model.CreateSubscriptionInput) (*model.SubscriptionStatus, error) {
	subscription_status := &model.SubscriptionStatus{}

	validation_error := validateInput(&input)

	if validation_error != nil {
		return subscription_status, validation_error
	}

	new_account := &models.Account{
		Address:      input.Address,
		Email:        input.Email,
		CurrencyCode: input.CurrencyCode.String(),
	}

	hook, hook_error := blockcypher.SetupHookOnAddress(new_account.Address, new_account.CurrencyCode)

	if hook_error != nil {
		return subscription_status, hook_error
	}

	new_account.HookId = hook.ID

	save_result := new_account.Create(database.DB)

	if save_result.Error != nil {
		blockcypher.DeleteHookOnAddress(hook.ID, new_account.CurrencyCode)
		return subscription_status, save_result.Error
	}

	subscription_status.IsSubscribed = true

	return subscription_status, nil
}

func CancelSubscription(ctx context.Context, input model.CancelSubscriptionInput) (*model.SubscriptionStatus, error) {
	subscription_status := &model.SubscriptionStatus{}

	account := models.Account{
		Address: input.Address,
		Email:   input.Email,
	}

	lookup_result := account.Get(database.DB)

	if lookup_result.Error != nil {
		if lookup_result.Error == gorm.ErrRecordNotFound {
			return subscription_status, errors.New("not exist")
		}

		return subscription_status, lookup_result.Error
	}

	delete_result := account.Delete(database.DB)

	if delete_result.Error != nil {
		return subscription_status, delete_result.Error
	}

	return subscription_status, nil
}

func GetSubscriptionStatus(ctx context.Context, input model.GetStatusInput) (*model.SubscriptionStatus, error) {
	subscription_status := &model.SubscriptionStatus{}

	account := &models.Account{
		Address: input.Address,
		Email:   input.Email,
	}

	lookup_result := account.Get(database.DB)

	if lookup_result.Error != nil {
		if lookup_result.Error == gorm.ErrRecordNotFound {
			subscription_status.IsSubscribed = false
			return subscription_status, nil
		}
		return subscription_status, lookup_result.Error
	} else {
		subscription_status.IsSubscribed = true
		return subscription_status, nil
	}
}
