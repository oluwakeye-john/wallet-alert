package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/blockcypher/gobcy"
	"github.com/oluwakeye-john/wallet-alert/blockcypher"
	"github.com/oluwakeye-john/wallet-alert/database"
	"github.com/oluwakeye-john/wallet-alert/models"
)

func BlockCypherHook(w http.ResponseWriter, r *http.Request) {
	defer json.NewEncoder(w).Encode("200")

	tx := gobcy.TX{}

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

	for x, i := range accounts {
		log.Println("User ", x+1, ": ", i.Email)
		balance, err := blockcypher.GetAddressBalance(i.Address, i.CurrencyCode)

		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Balance %f", balance)
	}

}
