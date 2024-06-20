package kafka

import (
	"context"
	"encoding/json"
	"log"

	k "github.com/segmentio/kafka-go"
	"studyum_notifications/internal/use/controllers/controller"
	"studyum_notifications/internal/use/controllers/controller/dto/requests"
)

type Kafka interface {
	Run() error

	HandleMessage(ctx context.Context, message []byte)
}

type kafka struct {
	kafka      *k.Reader
	controller controller.Controller
}

func NewKafka(k *k.Reader, controller controller.Controller) Kafka {
	return &kafka{
		kafka:      k,
		controller: controller,
	}
}

func (k *kafka) HandleMessage(ctx context.Context, message []byte) {
	var request requests.Message
	if err := json.Unmarshal(message, &request); err != nil {
		log.Printf("error unmarshalling message: %v", err)
		return
	}

	err := k.controller.SendMessage(ctx, request)
	if err != nil {
		log.Printf("error sending message: %v", err)
		return
	}
}
