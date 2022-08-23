package xkafka

import (
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/wahidx/kafkax/config"
)

func Ping() {
	conn, err := kafka.Dial("tcp", config.KAFKA_HOST[0])
	if err != nil {
		fmt.Println("Failed to connect to Kafka\n", err)
		return
	}

	fmt.Println("Kafka connected")
	conn.Close()
}

func Publish(topic string) {
	fmt.Println("Publishing to", topic)
}

func Listen(topic string) {
	fmt.Println("Listening to", topic)
}
