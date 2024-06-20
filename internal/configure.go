package internal

import (
	"studyum_notifications/internal/config"
	"studyum_notifications/internal/use"
	"studyum_notifications/internal/use/controllers"
	"studyum_notifications/internal/use/handlers"
)

func Configure() (*handlers.Handlers, *controllers.Controllers, error) {
	kafka, err := config.NotificationsKafka(config.Config.NotificationsKafka)
	if err != nil {
		return nil, nil, err
	}

	smtp, err := config.SMTP(config.Config.Smtp)
	if err != nil {
		return nil, nil, err
	}

	hndl, ctrl, err := use.New(kafka, smtp, config.Config.Smtp.From, nil)
	if err != nil {
		return nil, nil, err
	}

	return hndl, ctrl, nil
}
