package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oluwakeye-john/wallet-alert/blockcypher"
	"github.com/oluwakeye-john/wallet-alert/database"
	"github.com/oluwakeye-john/wallet-alert/models"
)

func BlockCypherHook(w http.ResponseWriter, r *http.Request) {
	// acknowledge webhook immediately
	json.NewEncoder(w).Encode("200")

	tx := blockcypher.NewTransaction()

	error := json.NewDecoder(r.Body).Decode(&tx)

	if error != nil {
		log.Println("error", error)
		return
	}

	accounts := []models.Account{}

	address := tx.Addresses[0]

	log.Println("Amount: ", tx.Total)
	log.Println("Address: ", address)

	database.DB.Find(&accounts, "address = ?", tx.Addresses[0])

	log.Println("Addresses: ", len(accounts))

	for x, account := range accounts {
		log.Println("User ", x+1, ": ", account.Email)

		if err := account.IncrementTransactionCount(database.DB).Error; err != nil {
			log.Println(err)
			continue
		}

		// // balance, balance_error := blockcypher.GetAddressBalance(account.Address, account.CurrencyCode)

		// // if balance_error != nil {
		// // 	log.Println(balance_error)
		// // 	return
		// // }

		// log.Printf("Balance %f", balance)
	}

}
