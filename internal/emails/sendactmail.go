package emails

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

// https://cloud.google.com/appengine/docs/standard/go111/mail/sending-receiving-with-mail-api?hl=pt-br

func SendActivationEmail(activation_url string) error {
	log.Printf("url: %v\n", activation_url)
	email, _ := os.LookupEnv("SENDER_EMAIL")
	pswd, _ := os.LookupEnv("SENDER_PSWD")

	// Sender data.
	from := email
	password := strings.ReplaceAll(pswd, " ", "")

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: Activation Email\r\n"+
			"\r\n"+
			"This is a test email message.\nUrl: %s",
		email,
		activation_url))

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Printf("Error sending email to %s\nError: %s", email, err)
		return err
	}
	fmt.Println("Email Sent Successfully!")
	return nil
}
