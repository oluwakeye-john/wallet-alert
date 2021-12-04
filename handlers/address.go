package handlers

import (
	"log"

	"github.com/blockcypher/gobcy"
	"github.com/oluwakeye-john/wallet-alert/blockcypher"
	"github.com/oluwakeye-john/wallet-alert/config"
	"github.com/oluwakeye-john/wallet-alert/graph/model"
)

func CreateTestAddress() (*model.Key, error) {
	bc := gobcy.API{}
	bc.Chain = "test"
	bc.Coin = "bcy"

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

func FundTestAddress(address string) (*model.Transaction, error) {
	bc := gobcy.API{}
	bc.Chain = "test"
	bc.Coin = "bcy"
	bc.Token = config.MustGetEnv("BLOCKCYPHER_KEY")

	transaction := &model.Transaction{}

	key := gobcy.AddrKeychain{
		Address: address,
	}

	txhash, err := bc.Faucet(key, 1e5)

	if err != nil {
		return transaction, err
	}

	transaction.Txhash = txhash

	return transaction, nil
}

func DeleteAddressHook(input model.DeleteHookInput) (bool, error) {
	err := blockcypher.DeleteHookOnAddress(input.HookID, string(input.CurrencyCode))

	if err != nil {
		return false, err
	}

	return true, nil
}
