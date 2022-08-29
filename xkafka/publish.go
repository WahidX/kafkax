package xkafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
	"github.com/wahidx/kafkax/config"
)

func Publish(topic, message, key string, headerMap map[string]string) {
	writer := &kafka.Writer{
		Addr:  kafka.TCP(config.KAFKA_HOST...),
		Topic: topic,
	}

	defer func() {
		if err := writer.Close(); err != nil {
			fmt.Println("Failed to close writer\n", "error:", err)
		}
	}()

	headers := []protocol.Header{}
	for k, v := range headerMap {
		headers = append(headers, protocol.Header{
			Key:   k,
			Value: []byte(v),
		})
	}

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:     []byte(key),
			Value:   []byte(message),
			Headers: headers,
		},
	)

	if err != nil {
		fmt.Println("Failed to publish message\n", err)
		return
	}
}
