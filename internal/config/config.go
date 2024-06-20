package config

import (
	"github.com/stdyum/api-common/env"
)

var Config Model

type Model struct {
	NotificationsKafka NotificationsKafkaConfig `env:"KAFKA_NOTIFICATIONS"`
	Smtp               SMTPConfig               `env:"SMTP"`
}

func init() {
	err := env.Fill(&Config)
	if err != nil {
		panic("cannot fill config: " + err.Error())
	}
}
