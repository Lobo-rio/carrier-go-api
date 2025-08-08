package mail

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(to string, subject string, body string) error {
	fmt.Println("Sending email...")

	dSendEmail := gomail.NewDialer(os.Getenv("GMAIL_SMTP"), 587, os.Getenv("GMAIL_USER"), os.Getenv("GMAIL_PASSWORD"))

	message := gomail.NewMessage()

	message.SetHeader("From", os.Getenv("GMAIL_USER"))
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

    return dSendEmail.DialAndSend(message)
}