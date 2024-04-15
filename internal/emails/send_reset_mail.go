package emails

import (
	"log"
	"net/smtp"
	"os"
)

// https://cloud.google.com/appengine/docs/standard/go111/mail/sending-receiving-with-mail-api?hl=pt-br
// https://mailtrap.io/blog/golang-send-email/

func SendResetEmail(email string, resetUrl string) error {
	from, _ := os.LookupEnv("SENDER_EMAIL")
	password, _ := os.LookupEnv("SENDER_PSWD")
	to := email

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: GoBot - Password reset\n\n" +
		"Navigate to the url below to reset your password\n" + resetUrl

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Email successfully sent to %s", to)
	return nil
}
