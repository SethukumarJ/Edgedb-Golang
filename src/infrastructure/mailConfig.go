package infrastructure

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type MailConfig interface {
	SendMail(to string, message []byte) error
}

type mailConfig struct{}

func NewMailConfig() MailConfig {
	return &mailConfig{}
}

func (c *mailConfig) SendMail(to string, message []byte) error {

	fmt.Printf("\n\nemail :  %v\n\n", to)
	log.Println("Email Id to send message : ", to)
	userName := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	auth := smtp.PlainAuth("", userName, password, smtpHost)
	fmt.Printf("\n\nauth : %v\n\n", auth)

	// sending email
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, userName, []string{to}, message)

}
