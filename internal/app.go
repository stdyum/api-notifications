package internal

import (
	"log"

	"studyum_notifications/internal/use/handlers/kafka"
)

func App() kafka.Kafka {
	routes, _, err := Configure()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return routes.Kafka
}
