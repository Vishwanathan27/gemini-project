package services

import (
	"fmt"
	"net/smtp"
	"os"
)

// MailSender contains the configuration for sending an email
type MailSender struct {
	Host     string
	Port     string
	Username string
	Password string
	Auth     smtp.Auth
}

// NewMailSender initializes a new MailSender with settings from environment variables
func NewMailSender() *MailSender {
	username := os.Getenv("MICROSOFT_EMAIL")
	password := os.Getenv("MICROSOFT_PASSWORD")

	host := "smtp-mail.outlook.com"
	port := "587"

	auth := smtp.PlainAuth("", username, password, host)

	return &MailSender{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Auth:     auth,
	}
}

// SendEmail sends an email using the Microsoft SMTP server
func (ms *MailSender) SendEmail(to []string, subject, body string) error {
	msg := []byte("From: " + ms.Username + "\r\n" +
		"To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body)

	serverAddr := fmt.Sprintf("%s:%s", ms.Host, ms.Port)
	err := smtp.SendMail(serverAddr, ms.Auth, ms.Username, to, msg)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	fmt.Println("Email sent successfully")
	return nil
}
