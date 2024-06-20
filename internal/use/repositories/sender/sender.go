package sender

import (
	"context"

	"gopkg.in/gomail.v2"
)

type Sender interface {
	Send(ctx context.Context, message Message) error
}

type sender struct {
	smtp *gomail.Dialer
	from string
}

func NewSender(smtp *gomail.Dialer, from string) Sender {
	return &sender{
		smtp: smtp,
		from: from,
	}
}

func (s *sender) Send(_ context.Context, message Message) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", message.To)
	m.SetHeader("Subject", message.Subject)
	m.SetBody("text/html", message.Content)

	return s.smtp.DialAndSend(m)
}
