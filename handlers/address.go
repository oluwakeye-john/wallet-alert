package handlers

import (
	"log"

	"github.com/oluwakeye-john/wallet-alert/blockcypher"
	"github.com/oluwakeye-john/wallet-alert/currencies"
	"github.com/oluwakeye-john/wallet-alert/customerrors"
	"github.com/oluwakeye-john/wallet-alert/graph/model"
	"github.com/oluwakeye-john/wallet-alert/models"
	"gorm.io/gorm"
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

func DeleteAddressWithNoAccount(db *gorm.DB, addr *models.Address) error {
	result := db.Preload("Accounts").First(addr, "id = ?", addr.ID)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Number of accounts", len(addr.Accounts))
	if len(addr.Accounts) == 0 {
		if error := blockcypher.DeleteAddressTransactionHook(addr.HookId, addr.CurrencyCode); error != nil {
			return error
		}

		if error := db.Delete(addr, "id = ?", addr.ID).Error; error != nil {
			return error
		}

		return nil
	}
	log.Println("Contains other accounts")
	return nil
}
