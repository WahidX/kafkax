/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wahidx/kafkax/config"
	"github.com/wahidx/kafkax/utils"
	"github.com/wahidx/kafkax/xkafka"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Test connectivity with a broker",

	Run: func(cmd *cobra.Command, args []string) {
		broker, err := cmd.Flags().GetString("broker")
		if err != nil {
			fmt.Println("No broker mentioned, using default: localhost:9092")
		} else {
			config.KAFKA_HOST = utils.ReadBrokers(broker)
		}

		xkafka.Ping()
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	rootCmd.PersistentFlags().StringP("broker", "b", "", "Kafka broker (comma separated)")
}
