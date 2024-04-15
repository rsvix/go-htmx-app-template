package emails

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// https://cloud.google.com/appengine/docs/standard/go111/mail/sending-receiving-with-mail-api?hl=pt-br

func SendActivationEmail(email string, activation_url string) error {
	log.Printf("url: %v\n", activation_url)

	// Sender data.
	from, _ := os.LookupEnv("SENDER_EMAIL")
	password, _ := os.LookupEnv("SENDER_PSWD")
	to := email

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpHostPort := "smtp.gmail.com:587"

	body := activation_url

	// Message.
	message := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHostPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		fmt.Printf("Error sending email to %s\nError: %s", email, err)
		return err
	}
	fmt.Printf("Email Sent Successfully to %s!", email)
	return nil
}
