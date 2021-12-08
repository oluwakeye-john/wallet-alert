package mail

import (
	"log"

	"github.com/oluwakeye-john/wallet-alert/models"
)

func SendNewSubscriberMail(account *models.Account) {
	mail_data := MailData{
		TemplateName: "new_subscriber",
		To:           account.Email,
		Subject:      "Welcome to Wallet-Alert",
		Data: struct {
			Email string
		}{
			Email: account.Email,
		},
	}

	err := SendMail(mail_data)
	if err != nil {
		log.Println(err)
	}
}
