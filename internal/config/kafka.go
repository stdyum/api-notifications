package config

import (
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

type NotificationsKafkaConfig struct {
	URL      string `env:"URL"`
	Topic    string `env:"TOPIC"`
	Group    string `env:"GROUP"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
}

func NotificationsKafka(config NotificationsKafkaConfig) (*kafka.ReaderConfig, error) {
	return &kafka.ReaderConfig{
		Brokers: []string{config.URL},
		Topic:   config.Topic,
		GroupID: config.Group,
		Dialer: &kafka.Dialer{
			Timeout:   10 * time.Second,
			DualStack: true,
			SASLMechanism: plain.Mechanism{
				Username: config.User,
				Password: config.Password,
			},
		},
	}, nil
}
