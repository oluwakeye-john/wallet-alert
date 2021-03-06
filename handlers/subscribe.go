package handlers

import (
	"context"
	"strings"

	"github.com/oluwakeye-john/wallet-alert/currencies"
	"github.com/oluwakeye-john/wallet-alert/customerrors"
	"github.com/oluwakeye-john/wallet-alert/database"
	"github.com/oluwakeye-john/wallet-alert/graph/model"
	"github.com/oluwakeye-john/wallet-alert/mail"
	"github.com/oluwakeye-john/wallet-alert/models"
	"github.com/oluwakeye-john/wallet-alert/utils/validators"
	"gorm.io/gorm"
)

func validateInput(input *model.CreateSubscriptionInput) error {
	is_email_valid := validators.IsEmailValid(input.Email)

	if !is_email_valid {
		return customerrors.InvalidEmail()
	}

	return nil
}

func CreateSubscription(ctx context.Context, input model.CreateSubscriptionInput) (*model.SubscriptionStatus, error) {
	subscription_status := &model.SubscriptionStatus{}

	validation_error := validateInput(&input)

	if validation_error != nil {
		return subscription_status, validation_error
	}

	currency, currency_error := currencies.GetCurrencyFromCode(string(input.CurrencyCode))

	if currency_error != nil {
		return subscription_status, currency_error
	}

	is_address_valid := currency.IsAddressValid(input.Address)

	if !is_address_valid {
		return subscription_status, customerrors.InvalidAddress()
	}

	account := &models.Account{
		Email:        input.Email,
		Address:      input.Address,
		CurrencyCode: currency.Code,
	}

	if err := account.Create(database.DB).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return subscription_status, customerrors.AlreadySubscribed()
		}
		return subscription_status, err
	}

	go account.SetupHook(database.DB)
	go mail.SendNewSubscriberMail(account)

	subscription_status.IsSubscribed = true
	subscription_status.Address = account.Address
	return subscription_status, nil
}

func CancelSubscription(ctx context.Context, input model.CancelSubscriptionInput) (*model.SubscriptionStatus, error) {
	subscription_status := &model.SubscriptionStatus{}

	account := models.Account{
		Email: input.Email,
	}

	if err := account.Delete(database.DB).Error; err != nil {
		return subscription_status, nil
	}

	go account.DeleteHook(database.DB)

	return subscription_status, nil
}

func GetSubscriptionStatus(ctx context.Context, input model.GetStatusInput) (*model.SubscriptionStatus, error) {
	subscription_status := &model.SubscriptionStatus{}

	account := models.Account{
		Email: input.Email,
	}

	if err := account.Get(database.DB).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return subscription_status, nil
		}
		return subscription_status, err
	}

	subscription_status.IsSubscribed = true
	subscription_status.Address = account.Address

	return subscription_status, nil
}
