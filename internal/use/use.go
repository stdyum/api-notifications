package use

import (
	"github.com/segmentio/kafka-go"
	"gopkg.in/gomail.v2"
	"studyum_notifications/internal/use/controllers"
	"studyum_notifications/internal/use/errors"
	"studyum_notifications/internal/use/handlers"
	"studyum_notifications/internal/use/repositories"
)

func New(kafkaConfig *kafka.ReaderConfig, smtp *gomail.Dialer, emailsFrom string, templates any) (*handlers.Handlers, *controllers.Controllers, error) {
	repos := repositories.New(smtp, emailsFrom, templates)

	ctrl := controllers.New(repos.Sender, repos.Templates)

	errors.Register()

	kafkaReader := kafka.NewReader(*kafkaConfig)
	h := handlers.New(kafkaReader, ctrl.Controller)

	return h, ctrl, nil
}
