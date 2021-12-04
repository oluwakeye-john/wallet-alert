package blockcypher

import (
	"log"
	"strings"

	"github.com/oluwakeye-john/wallet-alert/config"
	"github.com/oluwakeye-john/wallet-alert/utils"

	"github.com/blockcypher/gobcy"
)

func SetupHookOnAddress(address string, currency_code string) (gobcy.Hook, error) {
	log.Print("Configuring webhook...")
	defer log.Println("Done")

	coin, coin_error := utils.GetCoin(currency_code)
	chain := utils.GetChain(currency_code)

	if coin_error != nil {
		return gobcy.Hook{}, coin_error
	}

	bc := gobcy.API{}
	bc.Coin = strings.ToLower(coin)
	bc.Chain = chain
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

func DeleteHookOnAddress(hook_id string, currency_code string) error {
	log.Print("Destroying webhook...")
	defer log.Println("Done")

	coin, coin_error := utils.GetCoin(currency_code)
	chain := utils.GetChain(currency_code)

	if coin_error != nil {
		return coin_error
	}

	bc := gobcy.API{}
	bc.Coin = strings.ToLower(coin)
	bc.Chain = chain
	bc.Token = config.MustGetEnv("BLOCKCYPHER_KEY")

	err := bc.DeleteHook(hook_id)

	if err != nil {
		return err
	}

	return nil
}
