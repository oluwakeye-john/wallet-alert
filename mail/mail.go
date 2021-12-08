package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/oluwakeye-john/wallet-alert/config"
)

type MailData struct {
	TemplateName string
	To           string
	Subject      string
	Data         interface{}
}

func SendMail(mail_data MailData) error {
	log.Printf("Sending %s mail to %s", mail_data.TemplateName, mail_data.To)
	tmp, err := template.ParseFiles(fmt.Sprintf("mail/templates/%s.html", mail_data.TemplateName))

	if err != nil {
		return err
	}

	buf := bytes.Buffer{}

	err = tmp.Execute(&buf, mail_data.Data)

	if err != nil {
		return err
	}

	host := config.GetEnv("MAIL_HOST")
	port := config.GetEnv("MAIL_PORT")

	mail_from := config.GetEnv("MAIL_FROM")
	mail_username := config.GetEnv("MAIL_USERNAME")
	mail_password := config.GetEnv("MAIL_PASSWORD")

	mail := email.NewEmail()
	mail.From = fmt.Sprintf("Wallet Alert <%s>", mail_from)
	mail.To = []string{mail_data.To}
	mail.Subject = mail_data.Subject
	mail.HTML = buf.Bytes()
	// mail.ReplyTo = []string{""}

	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", mail_username, mail_password, host)

	err = mail.Send(addr, auth)

	if err != nil {
		return err
	}

	log.Println("Mail sent")

	return nil
}
