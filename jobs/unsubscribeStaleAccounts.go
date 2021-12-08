package jobs

import (
	"log"
	"time"

	"github.com/oluwakeye-john/wallet-alert/database"
	"github.com/oluwakeye-john/wallet-alert/models"
)

func UnsubscribeStaleAccounts() {
	accounts := []models.Account{}

	limit := time.Now().AddDate(0, -3, 0)

	if err := database.DB.Find(&accounts, "updated_at < ?", limit); err != nil {
		log.Println("Error fetching account")
		return
	}

	for _, a := range accounts {
		go func(account models.Account) {
			if err := account.Delete(database.DB).Error; err != nil {
				account.DeleteHook(database.DB)
			}
		}(a)

		// 1 second delay
		time.Sleep(time.Second)
	}

}
