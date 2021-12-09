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
			Email   string
			Address string
		}{
			Email:   account.Email,
			Address: account.Address,
		},
	}

	err := SendMail(mail_data)
	if err != nil {
		log.Println(err)
	}
}
