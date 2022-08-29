package utils

import (
	"strings"

	"github.com/wahidx/kafkax/config"
)

func SetBroker(broker string) {
	brokers := strings.Split(broker, ",")
	config.KAFKA_HOST = brokers
}
