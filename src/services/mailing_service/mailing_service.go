package mailing_service

import (
	"fmt"
	"net/smtp"
	"strings"

	"delivery_app_api.mmedic.com/m/v2/src/utils/env_utils"
)

type Email struct {
	sender   string
	receiver []string
	subject  string
	body     string
}

func CreateEmail(receiver, subject, body string) *Email {
	return &Email{
		sender:   env_utils.GetEnvVar("MAIL_SERVICE_EMAIL"),
		receiver: strings.Split(receiver, " "),
		subject:  subject,
		body:     body,
	}
}

func (e Email) BuildMessage() string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", e.sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(e.receiver, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", e.subject)
	msg += fmt.Sprintf("\r\n%s\r\n", e.body)

	return msg
}

func (e *Email) SendMail() error {
	from := env_utils.GetEnvVar("MAIL_SERVICE_EMAIL")
	password := env_utils.GetEnvVar("MAIL_SERVICE_PASS")

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	message := e.BuildMessage()

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, e.receiver, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
