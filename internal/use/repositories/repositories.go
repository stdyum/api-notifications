package repositories

import (
	"gopkg.in/gomail.v2"
	"studyum_notifications/internal/use/repositories/sender"
	"studyum_notifications/internal/use/repositories/templates"
)

type Repositories struct {
	Sender    sender.Sender
	Templates templates.Templates
}

func New(smtp *gomail.Dialer, from string, a any) Repositories {
	return Repositories{
		Sender:    sender.NewSender(smtp, from),
		Templates: templates.NewTemplates(a),
	}
}
