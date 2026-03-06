package util

import (
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(toEmail string, subject string, body string) error {

	from := mail.NewEmail("Hotel API", os.Getenv("EMAIL_SENDER"))
	to := mail.NewEmail("", toEmail)

	message := mail.NewSingleEmail(from, subject, to, body, body)

	client := sendgrid.NewSendClient(os.Getenv("API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		return err
	}

	println("Sendgrid status:", response.StatusCode)
	println("Sendgrid body:", response.Body)

	return nil
}
