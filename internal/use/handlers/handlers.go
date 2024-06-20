package handlers

import (
	k "github.com/segmentio/kafka-go"
	"studyum_notifications/internal/use/controllers/controller"
	"studyum_notifications/internal/use/handlers/kafka"
)

type Handlers struct {
	Kafka kafka.Kafka
}

func New(k *k.Reader, controller controller.Controller) *Handlers {
	return &Handlers{
		Kafka: kafka.NewKafka(k, controller),
	}
}
