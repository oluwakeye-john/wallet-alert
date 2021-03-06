package mail

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/blockcypher/gobcy"
	"github.com/oluwakeye-john/wallet-alert/currencies"
	"github.com/oluwakeye-john/wallet-alert/models"
)

func SendTransactionMail(account *models.Account, tx *gobcy.TX, balance float64) {
	currency, currency_error := currencies.GetCurrencyFromCode(account.CurrencyCode)

	if currency_error != nil {
		log.Println(currency_error)
		return
	}

	tx_amount, _ := strconv.ParseFloat(tx.Total.String(), 64)

	location, _ := time.LoadLocation("UTC")

	mail_data := MailData{
		TemplateName: "new_transaction",
		To:           account.Email,
		Subject:      "Transaction Notification",
		Data: struct {
			AmountInString  string
			BalanceInString string
			Address         string
			CurrencyCode    string
			Type            string
			Time            string
			Email           string
		}{
			Time:            tx.Confirmed.In(location).Format(time.RFC850),
			BalanceInString: fmt.Sprintf("%f%s", balance/100000000, currency.Code),
			AmountInString:  fmt.Sprintf("%f%s", tx_amount/100000000, currency.Code),
			Address:         account.Address,
			CurrencyCode:    currency.Code,
			Type:            "incoming",
			Email:           account.Email,
		},
	}

	err := SendMail(mail_data)
	if err != nil {
		log.Println(err)
	}
}
