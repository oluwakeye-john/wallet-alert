package blockcypher

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/oluwakeye-john/wallet-alert/config"
	"github.com/oluwakeye-john/wallet-alert/currencies"
	"github.com/oluwakeye-john/wallet-alert/customerrors"
	"github.com/oluwakeye-john/wallet-alert/graph/model"

	"github.com/blockcypher/gobcy"
)

func initGobcyApi(currency currencies.Currency) *gobcy.API {
	bc := gobcy.API{}
	bc.Chain = currency.Chain
	bc.Coin = strings.ToLower(currency.Code)
	bc.Token = config.MustGetEnv("BLOCKCYPHER_API_KEY")

	return &bc
}

func SetupAddressTransactionHook(address string, currency_code string) (gobcy.Hook, error) {
	log.Printf("Configuring webhook for %s...", address)
	defer log.Println("Done")

	currency, currency_error := currencies.GetCurrencyFromCode(currency_code)

	if currency_error != nil {
		log.Println(currency_error)
		return gobcy.Hook{}, currency_error
	}

	bc := initGobcyApi(currency)

	hook := gobcy.Hook{
		Event:   "confirmed-tx",
		Address: address,
		URL:     fmt.Sprintf("%s/cb", config.MustGetEnv("SITE_DOMAIN")),
	}

	log.Println(hook)

	hook, err := bc.CreateHook(hook)

	if err != nil {
		log.Println(err)
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

	bc := initGobcyApi(currency)

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

	bc := initGobcyApi(currency)

	addr, err := bc.GetAddrBal(address, nil)

	if err != nil {
		return 0, err
	}

	balance, _ := new(big.Float).SetInt(&addr.Balance).Float64()
	return balance, nil
}

func CreateTestAddress() (*model.Address, error) {
	currency := currencies.TestCurrency

	bc := initGobcyApi(currency)

	var err error

	log.Println("Generating test address")
	keys, err := bc.GenAddrKeychain()

	if err != nil {
		return &model.Address{}, err
	}

	address := &model.Address{
		Address:      keys.Address,
		PublicKey:    keys.Public,
		PrivateKey:   keys.Private,
		CurrencyCode: model.CurrencyCode(currency.Code),
		ExplorerURL:  currency.ExplorerAddrUrl(keys.Address),
	}

	return address, nil
}

func FundTestAddress(address string, amount float64) (*model.Transaction, error) {
	currency := currencies.TestCurrency

	transaction := &model.Transaction{
		Amount: amount,
	}

	bc := initGobcyApi(currency)

	is_address_valid := currencies.TestCurrency.IsAddressValid(address)

	if !is_address_valid {
		return transaction, customerrors.InvalidAddress()
	}

	key := gobcy.AddrKeychain{
		Address: address,
	}

	txhash, err := bc.Faucet(key, int(transaction.Amount))

	if err != nil {
		return transaction, err
	}

	transaction.Txhash = txhash
	transaction.ExplorerURL = currency.ExplorerTxUrl(txhash)

	return transaction, nil
}

func NewTransaction() *gobcy.TX {
	return &gobcy.TX{}
}
