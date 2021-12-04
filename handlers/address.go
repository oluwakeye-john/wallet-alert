package handlers

import (
	"log"
	"strings"

	"github.com/blockcypher/gobcy"
	"github.com/oluwakeye-john/wallet-alert/blockcypher"
	"github.com/oluwakeye-john/wallet-alert/config"
	"github.com/oluwakeye-john/wallet-alert/currencies"
	"github.com/oluwakeye-john/wallet-alert/customerrors"
	"github.com/oluwakeye-john/wallet-alert/graph/model"
)

func CreateTestAddress() (*model.Key, error) {
	currency := currencies.TestCurrency

	bc := gobcy.API{}
	bc.Chain = currency.Chain
	bc.Coin = currency.CodeInLowerCase()

	bc.Token = config.MustGetEnv("BLOCKCYPHER_KEY")

	var err error

	log.Println("Generating test address")
	keys, err := bc.GenAddrKeychain()

	if err != nil {
		return &model.Key{}, err
	}

	address := &model.Key{
		Address:      keys.Address,
		PublicKey:    keys.Public,
		PrivateKey:   keys.Private,
		CurrencyCode: model.CurrencyCode("BCY"),
	}

	return address, nil

}

func FundTestAddress(input model.FundTestAddressInput) (*model.Transaction, error) {
	currency := currencies.TestCurrency

	transaction := &model.Transaction{
		Amount: input.Amount,
	}

	bc := gobcy.API{}
	bc.Chain = currency.Chain
	bc.Coin = strings.ToLower(currency.Code)
	bc.Token = config.MustGetEnv("BLOCKCYPHER_KEY")

	is_address_valid := currencies.TestCurrency.IsValid(input.Address)

	if !is_address_valid {
		return transaction, customerrors.InvalidAddress()
	}

	key := gobcy.AddrKeychain{
		Address: input.Address,
	}

	txhash, err := bc.Faucet(key, int(transaction.Amount))

	if err != nil {
		return transaction, err
	}

	transaction.Txhash = txhash

	return transaction, nil
}

func DeleteAddressHook(input model.DeleteHookInput) (bool, error) {
	err := blockcypher.DeleteAddressTransactionHook(input.HookID, string(input.CurrencyCode))

	if err != nil {
		return false, err
	}

	return true, nil
}
