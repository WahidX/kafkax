package xkafka

import (
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/wahidx/kafkax/config"
)

func ListTopics() {
	if len(config.KAFKA_HOST) == 0 {
		fmt.Println("Broker is blank")
		return
	}

	conn, err := kafka.Dial("tcp", config.KAFKA_HOST[0])
	if err != nil {
		fmt.Println("Failed to connect to Kafka\n", err)
		return
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		fmt.Println("Failed to read partitions\n", err)
		return
	}

	m := map[string]bool{}
	for _, p := range partitions {
		m[p.Topic] = true
	}

	fmt.Println("Topics:")
	for k := range m {
		fmt.Println(k)
	}
}
