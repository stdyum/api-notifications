package controllers

import (
	"studyum_notifications/internal/use/controllers/controller"
	"studyum_notifications/internal/use/repositories/sender"
	"studyum_notifications/internal/use/repositories/templates"
)

type Controllers struct {
	Controller controller.Controller
}

func New(sender sender.Sender, templates templates.Templates) *Controllers {
	return &Controllers{
		Controller: controller.NewController(sender, templates),
	}
}
