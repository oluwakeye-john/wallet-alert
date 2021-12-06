package handlers

import (
	"github.com/oluwakeye-john/wallet-alert/blockcypher"
	"github.com/oluwakeye-john/wallet-alert/currencies"
	"github.com/oluwakeye-john/wallet-alert/customerrors"
	"github.com/oluwakeye-john/wallet-alert/graph/model"
)

func CreateTestAddress() (*model.Address, error) {
	return blockcypher.CreateTestAddress()
}

func FundTestAddress(input model.FundTestAddressInput) (*model.Transaction, error) {
	currency := currencies.TestCurrency

	if !currency.IsAddressValid(input.Address) {
		return &model.Transaction{}, customerrors.InvalidAddress()
	}

	return blockcypher.FundTestAddress(input.Address, input.Amount)
}

func DeleteAddressHook(input model.DeleteHookInput) (bool, error) {
	currency, currency_error := currencies.GetCurrencyFromCode(string(input.CurrencyCode))

	if currency_error != nil {
		return false, currency_error
	}

	err := blockcypher.DeleteAddressTransactionHook(input.HookID, currency.Code)

	if err != nil {
		return false, err
	}

	return true, nil
}
