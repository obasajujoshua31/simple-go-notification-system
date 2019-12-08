package services

import (
	"errors"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"simple-go-notification-system/config"
)

type Message struct {
	Subject   string
	ToEmail   string
	PlainText string
	FromName  string
	FromEmail string
	ToName    string
	HTMLContent string
}

type MailClient struct {
	Client *sendgrid.Client
}

func NewMailClient(config config.AppConfig) MailClient {
	client := sendgrid.NewSendClient(config.SendGridAPI)
	return MailClient{
		Client: client,
	}
}

func (cl *MailClient) SendEmailMessage(message *mail.SGMailV3) error {
	resp, err := cl.Client.Send(message)

	if err != nil {

		return err
	}

	if resp.StatusCode >=300 {
		return errors.New("unable to send Email")
	}
	fmt.Println("Response", resp.StatusCode, resp.Body)
	return nil
}


func  NewEmailMessage(message *Message) *mail.SGMailV3 {
	sender := mail.NewEmail(message.FromName, message.FromEmail)
	to := mail.NewEmail(message.ToName, message.ToEmail)

	ms := mail.NewSingleEmail(sender, message.Subject, to, message.PlainText, message.HTMLContent )

	return ms
}
