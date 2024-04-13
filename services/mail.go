package services

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailSender struct {
	APIKey string
}

func NewMailSender() *MailSender {
	apiKey := os.Getenv("SENDGRID_API_KEY") // The SendGrid API Key must be set as an environment variable
	if apiKey == "" {
		log.Fatal("SENDGRID_API_KEY environment variable not set.")
	}
	return &MailSender{
		APIKey: apiKey,
	}
}

func (ms *MailSender) SendEmail(subject, plainTextContent, htmlContent string) error {
	senderEmail := os.Getenv("SENDER_EMAIL")
	receiverEmail := os.Getenv("RECEIVER_EMAIL")
	senderName := os.Getenv("SENDER_NAME")
	receiverName := os.Getenv("RECEIVER_NAME")
	if senderEmail == "" || receiverEmail == "" || senderName == "" || receiverName == "" {
		log.Println("One or more required environment variables are missing:")
		log.Printf("SENDER_EMAIL: %s, RECEIVER_EMAIL: %s, SENDER_NAME: %s, RECEIVER_NAME: %s\n",
			senderEmail, receiverEmail, senderName, receiverName)
		return nil
	}

	from := mail.NewEmail(senderName, senderEmail)
	to := mail.NewEmail(receiverName, receiverEmail)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(ms.APIKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println("Email Sent Successfully!")
		fmt.Println("Status Code: ", response.StatusCode)
		fmt.Println("Response", response)
		fmt.Println("Response Body: ", response.Body)
		fmt.Println("Response Headers: ", response.Headers)
		return nil
	}
}
