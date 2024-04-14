package emails

import (
	"fmt"
	"log"
	"net/smtp"
)

// https://cloud.google.com/appengine/docs/standard/go111/mail/sending-receiving-with-mail-api?hl=pt-br
// https://mailtrap.io/blog/golang-send-email/

func SendResetEmail(email string, activation_url string) error {
	log.Printf("url: %v\n", activation_url)
	log.Printf("Send email to: %s\n", email)

	// Sender data.
	from := "from@gmail.com"
	password := "<Email Password>"

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
			"This is a test email message.\n%s",
		email,
		activation_url))

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")
	return nil
}
