package xkafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/wahidx/kafkax/config"
)

func Publish(topic, message, key string) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(config.KAFKA_HOST...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer func() {
		if err := writer.Close(); err != nil {
			fmt.Println("Failed to close writer\n", "error:", err)
		}
	}()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(message),
		},
	)

	if err != nil {
		fmt.Println("Failed to publish message\n", err)
		return
	}

	fmt.Println("Message published!")
}
