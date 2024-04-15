package emails

import (
	"log"
	"net/smtp"
	"os"
)

// https://gist.github.com/jpillora/cb46d183eca0710d909a

func SendActivationEmail(email string, activationUrl string) error {
	from, _ := os.LookupEnv("SENDER_EMAIL")
	password, _ := os.LookupEnv("SENDER_PSWD")
	to := email

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: GoBot - Account ativation\n\n" +
		"Navigate to the url below to activate your account\n" + activationUrl

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Printf("Error sending email to %s\nError: %s", email, err)
		return err
	}
	log.Printf("Email successfully sent to %s", to)
	return nil
}
