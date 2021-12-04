package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/blockcypher/gobcy"
)

func BlockCypherHook(w http.ResponseWriter, r *http.Request) {
	hook := gobcy.Hook{}

	error := json.NewDecoder(r.Body).Decode(&hook)

	if error != nil {
		log.Println("error", error)
		return
	}

	log.Println(hook)

	json.NewEncoder(w).Encode("200")
}
