package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

	hook_address := tx.Addresses[0]

	total, _ := strconv.ParseFloat(tx.Total.String(), 64)

	log.Println("Amount: ", total)
	log.Println("Address: ", hook_address)

	address := models.Address{}

	if err := database.DB.First(&address, "address = ?", hook_address).Error; err != nil {
		log.Println(err)
		return
	}

	accounts := []models.Account{}

	if err := database.DB.Find(&accounts, "address_id = ?", address.ID).Error; err != nil {
		log.Println(err)
		return
	}

	log.Println("Address: ", address.Address)

	for x, account := range accounts {
		log.Println("User ", x+1, ": ", account.Email)

		if err := account.IncrementTransactionCount(database.DB); err != nil {
			log.Println(err)
			continue
		}
	}
}
