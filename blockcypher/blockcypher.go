package blockcypher

import (
	"log"
	"math/big"

	"github.com/oluwakeye-john/wallet-alert/config"
	"github.com/oluwakeye-john/wallet-alert/currencies"

	"github.com/blockcypher/gobcy"
)

func SetupAddressTransactionHook(address string, currency_code string) (gobcy.Hook, error) {
	log.Print("Configuring webhook...")
	defer log.Println("Done")

	currency, currency_error := currencies.GetCurrencyFromCode(currency_code)

	if currency_error != nil {
		return gobcy.Hook{}, currency_error
	}

	bc := gobcy.API{}
	bc.Coin = currency.CodeInLowerCase()
	bc.Chain = currency.Chain
	bc.Token = config.MustGetEnv("BLOCKCYPHER_KEY")

	hook, err := bc.CreateHook(gobcy.Hook{
		Event:   "confirmed-tx",
		Address: address,
		URL:     config.MustGetEnv("BLOCKCYPHER_WEBHOOK_URL"),
	})

	if err != nil {
		return gobcy.Hook{}, err
	}

	return hook, nil
}

func DeleteAddressTransactionHook(hook_id string, currency_code string) error {
	log.Print("Destroying webhook...")
	defer log.Println("Done")

	currency, currency_error := currencies.GetCurrencyFromCode(currency_code)

	if currency_error != nil {
		return currency_error
	}

	bc := gobcy.API{}
	bc.Coin = currency.CodeInLowerCase()
	bc.Chain = currency.Chain
	bc.Token = config.MustGetEnv("BLOCKCYPHER_KEY")

	err := bc.DeleteHook(hook_id)

	if err != nil {
		return err
	}

	return nil
}

func GetAddressBalance(address string, currency_code string) (float64, error) {
	log.Print("Fetching address balance...")
	defer log.Println("Done")

	currency, currency_error := currencies.GetCurrencyFromCode(currency_code)

	if currency_error != nil {
		return 0, currency_error
	}

	bc := gobcy.API{}
	bc.Coin = currency.CodeInLowerCase()
	bc.Chain = currency.Chain

	addr, err := bc.GetAddrBal(address, nil)

	if err != nil {
		return 0, err
	}

	balance, _ := new(big.Float).SetInt(&addr.Balance).Float64()
	return balance, nil
}
