package utils

import "strings"

func ReadBrokers(broker string) []string {
	return strings.Split(broker, ",")
}
