package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oluwakeye-john/wallet-alert/blockcypher"
	"github.com/oluwakeye-john/wallet-alert/database"
	"github.com/oluwakeye-john/wallet-alert/mail"
	"github.com/oluwakeye-john/wallet-alert/models"
)

func BlockCypherHook(w http.ResponseWriter, r *http.Request) {
	// acknowledge webhook immediately
	log.Println("Incoming webhook")

	tx := blockcypher.NewTransaction()
	error := json.NewDecoder(r.Body).Decode(&tx)

	if error != nil {
		log.Println(error)
		return
	}

	go func() {
		hook_address := tx.Addresses[0]

		account := models.Account{
			Address: hook_address,
		}

		if err := account.GetByAddress(database.DB).Error; err != nil {
			log.Println(err)
			return
		}

		balance, err := blockcypher.GetAddressBalance(account.Address, account.CurrencyCode)

		if err != nil {
			log.Println(err)
			return
		}

		mail.SendTransactionMail(&account, tx, balance)

		if err := account.IncrementTransactionCount(database.DB); err != nil {
			log.Println(err)
			return
		}

	}()

	w.Write([]byte(""))
}
