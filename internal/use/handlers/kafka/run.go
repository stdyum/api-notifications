package kafka

import (
	"context"
	"fmt"
)

func (k *kafka) Run() error {
	ctx := context.Background()
	for {
		message, err := k.kafka.ReadMessage(ctx)
		if err != nil {
			return err
		}

		k.HandleMessage(ctx, message.Value)
		fmt.Printf("message at offset %d\n", message.Offset)
	}
}
