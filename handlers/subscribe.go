package handlers

import (
	"context"
	"errors"

	"github.com/oluwakeye-john/wallet-alert/currencies"
	"github.com/oluwakeye-john/wallet-alert/customerrors"
	"github.com/oluwakeye-john/wallet-alert/database"
	"github.com/oluwakeye-john/wallet-alert/graph/model"
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

	// here

	address := models.Address{
		Address:      input.Address,
		CurrencyCode: currency.Code,
	}

	if err := address.GetOrCreate(database.DB).Error; err != nil {
		return subscription_status, err
	}

	account := &models.Account{
		Email:     input.Email,
		AddressId: address.ID,
	}

	lookup_result := account.Get(database.DB)

	if lookup_result.Error == nil {
		return subscription_status, customerrors.AlreadySubscribed()
	} else if lookup_result.Error == gorm.ErrRecordNotFound {
		if err := account.Create(database.DB).Error; err != nil {
			return subscription_status, err
		}

		subscription_status.IsSubscribed = true
		return subscription_status, nil
	} else {
		return subscription_status, lookup_result.Error
	}

}

func CancelSubscription(ctx context.Context, input model.CancelSubscriptionInput) (*model.SubscriptionStatus, error) {
	subscription_status := &model.SubscriptionStatus{}

	account := models.Account{
		Email: input.Email,
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
		Email: input.Email,
	}

	if err := account.Get(database.DB).Error; err != nil {
		return subscription_status, err
	}

	address := &models.Address{}
	address.ID = account.AddressId

	if err := address.Get(database.DB).Error; err != nil {
		return subscription_status, err
	}

	subscription_status.IsSubscribed = true
	subscription_status.Address = address.Address

	return subscription_status, nil

}
