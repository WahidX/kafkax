package utils

import (
	"fmt"
	"strings"
)

func ReadBrokers(broker string) []string {
	return strings.Split(broker, ",")
}

func ReadHeaders(headers []string) map[string]string {
	header := map[string]string{}
	for _, h := range headers {
		val := strings.Split(h, "=")
		if len(val) != 2 {
			fmt.Println("Invalid header format\nShould be -H key=value\nFor multiple -H key1=val1 -H key2=val2")
			return nil
		}

		header[val[0]] = val[1]
	}

	return header
}
