package controller

import (
	"context"

	"studyum_notifications/internal/use/controllers/controller/dto/requests"
	"studyum_notifications/internal/use/repositories/sender"
	"studyum_notifications/internal/use/repositories/templates"
)

type Controller interface {
	SendMessage(ctx context.Context, message requests.Message) error
}

type controller struct {
	sender    sender.Sender
	templates templates.Templates
}

func NewController(sender sender.Sender, templates templates.Templates) Controller {
	return &controller{
		sender:    sender,
		templates: templates,
	}
}

func (c *controller) SendMessage(ctx context.Context, message requests.Message) error {
	template, err := c.templates.GetBodyByID(ctx, message.TemplateID, message.Data)
	if err != nil {
		return err
	}

	return c.sender.Send(ctx, sender.Message{
		To:      message.Email,
		Subject: template.Subject,
		Content: template.Content,
	})
}
