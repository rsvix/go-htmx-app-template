package emails

import (
	"log"
	"net/smtp"
	"os"
)

// https://gist.github.com/jpillora/cb46d183eca0710d909a

type mailParams struct {
	from     string
	password string
	smtpHost string
	smtpPort string
}

func DefaultParams() *mailParams {
	return &mailParams{
		from:     os.Getenv("SMTP_EMAIL"),
		password: os.Getenv("SMTP_PSWD"),
		smtpHost: "smtp.gmail.com",
		smtpPort: "587",
	}
}

func SendActivationMail(to string, tokenUrl string, p *mailParams) error {
	message := "From: " + p.from + "\n" +
		"To: " + to + "\n" +
		"Subject: GoBot - Account ativation\n\n" +
		"Navigate to the url below to activate your account\n" + tokenUrl

	auth := smtp.PlainAuth("", p.from, p.password, p.smtpHost)

	err := smtp.SendMail(p.smtpHost+":"+p.smtpPort, auth, p.from, []string{to}, []byte(message))
	if err != nil {
		return err
	}
	log.Printf("Email successfully sent to %s", to)
	return nil
}

func SendResetMail(to string, tokenUrl string, p *mailParams) error {
	message := "From: " + p.from + "\n" +
		"To: " + to + "\n" +
		"Subject: GoBot - Password reset\n\n" +
		"Navigate to the url below to reset your password\n" + tokenUrl

	auth := smtp.PlainAuth("", p.from, p.password, p.smtpHost)

	err := smtp.SendMail(p.smtpHost+":"+p.smtpPort, auth, p.from, []string{to}, []byte(message))
	if err != nil {
		return err
	}
	log.Printf("Email successfully sent to %s", to)
	return nil
}
