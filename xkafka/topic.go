package xkafka

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/topics"
	"github.com/wahidx/kafkax/config"
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

	m := map[string][]int{}
	for _, p := range partitions {
		if partitions := m[p.Topic]; partitions != nil {
			m[p.Topic] = append(m[p.Topic], p.ID)
		} else {
			m[p.Topic] = []int{p.ID}
		}
	}

	count := 1
	fmt.Println("Topics:")
	for topicName := range m {
		fmt.Println(strconv.Itoa(count) + ". " + topicName + " (partitions:" + fmt.Sprint(m[topicName]) + " )")
		count++
	}
}

func FindTopics(key string) {
	topics, err := topics.ListRe(context.Background(),
		&kafka.Client{Addr: kafka.TCP(config.KAFKA_HOST...)},
		regexp.MustCompile(key),
	)

	if err != nil {
		fmt.Println("Failed to search with key: ", key, "\nerorr:", err)
		return
	}

	for _, t := range topics {
		fmt.Println(t.Name)
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
