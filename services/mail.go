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
	from := mail.NewEmail("ViswaNewsLetter", "vish@getnada.com")
	to := mail.NewEmail("Viswanathan V", "vishhh27@outlook.com")

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(ms.APIKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println("Email Sent Successfully!")
		fmt.Println("Status Code: ", response.StatusCode)
		fmt.Println("Response Body: ", response.Body)
		fmt.Println("Response Headers: ", response.Headers)
		return nil
	}
}
