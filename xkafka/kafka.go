package xkafka

import (
	"fmt"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
	"github.com/wahidx/kafkax/config"
)

// getConnection returns the low level kafka connection object
func getConnection() (*kafka.Conn, error) {
	conn, err := kafka.Dial("tcp", config.KAFKA_HOST[0])
	if err != nil {
		fmt.Println("Failed to connect to Kafka\n", err)
		return nil, err
	}

	return conn, nil
}

// getControllerConnection returns the low level kafka connection object gotten through controllers
func getControllerConnection() (*kafka.Conn, error) {
	conn, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		fmt.Println("Failed to get controllers\n", err)
		return nil, err
	}

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		fmt.Println("Failed to get controller connection\n", err)
		return nil, err
	}

	return controllerConn, nil
}

func Ping() error {
	conn, err := getConnection()
	if err != nil {
		return err
	}

	fmt.Println("Kafka connected")
	conn.Close()
	return nil
}
