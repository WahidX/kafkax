package xkafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/wahidx/kafkax/config"
)

func Consume(opts ConsumerOptions) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     config.KAFKA_HOST,
		GroupID:     opts.GroupID,
		Topic:       opts.Topic,
		MinBytes:    5,
		MaxBytes:    10e6, // 10MB
		Partition:   opts.Partition,
		StartOffset: kafka.FirstOffset,
	})

	if opts.Offset != 0 {
		r.SetOffset(opts.Offset)
	}

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Failed to read message\n", err)
			break
		}

		printer(m, opts.IsJSON)
	}

	if err := r.Close(); err != nil {
		fmt.Println("Failed to close kafka reader\n", err)
	}
}

func printer(msg kafka.Message, isJSON bool) {
	if !isJSON {
		fmt.Println(string(msg.Value))
		return
	}

	var msgPayload any

	if string(msg.Value[0]) == "{" && string(msg.Value[len(msg.Value)-1]) == "}" {
		json.Unmarshal(msg.Value, &msgPayload) // nolint
	}

	// In case the message is not a valid json message then will be printed as it is
	if msgPayload == nil {
		msgPayload = string(msg.Value)
	}

	data := map[string]any{
		"topic":         msg.Topic,
		"timestamp":     msg.Time,
		"partition":     msg.Partition,
		"offset":        msg.Offset,
		"highWaterMark": msg.HighWaterMark,
		"key":           string(msg.Key),
		"data":          msgPayload,
		"headers":       msg.Headers,
	}

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Failed to format message in JSON\n", err)
		return
	}

	fmt.Println(string(b) + "\n")
}
