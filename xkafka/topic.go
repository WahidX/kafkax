package xkafka

import (
	"fmt"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func ListTopics() {
	conn, err := getConnection()
	if err != nil {
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

	count := 1
	fmt.Println("Topics:")
	for k := range m {
		fmt.Println(strconv.Itoa(count) + ". " + k)
		count++
	}
}

func CreateTopic(topic string, partition int, replica int) {
	conn, err := getControllerConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     partition,
			ReplicationFactor: replica,
		},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		fmt.Println("Failed to create topics\n", err)
		return
	}

	fmt.Println("Topic '" + topic + "' created successfully")
}

func DeleteTopic(topics ...string) {
	conn, err := getControllerConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	err = conn.DeleteTopics(topics...)
	if err != nil {
		fmt.Println("Failed to delete topic\n", err)
		return
	}

	fmt.Println("Topics deleted successfully")
}
