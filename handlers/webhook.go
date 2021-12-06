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
	log.Println("acknowledge")
	json.NewEncoder(w).Encode("200")

	tx := blockcypher.NewTransaction()
	error := json.NewDecoder(r.Body).Decode(&tx)

	if error != nil {
		log.Println(error)
		return
	}

	go func() {
		hook_address := tx.Addresses[0]

		address := models.Address{
			Address: hook_address,
		}

		type Result struct {
			models.Address
			models.Account
		}

		result := Result{}

		if err := database.DB.Table("addresses").Select("*").Joins("JOIN accounts ON accounts.address_id = addresses.id").First(&result, "addresses.address = ?", hook_address).Error; err != nil {
			log.Println(err)
			return
		}

		balance, err := blockcypher.GetAddressBalance(address.Address, address.CurrencyCode)

		if err != nil {
			log.Println(err)
			return
		}

		log.Println(address)

		for x, account := range address.Accounts {
			log.Println("User ", x+1, ": ", account.Email)

			mail.SendTransactionMail(&address, &account, tx, balance)

			if err := account.IncrementTransactionCount(database.DB); err != nil {
				log.Println(err)
				continue
			}
		}

	}()
}
