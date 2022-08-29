/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wahidx/kafkax/utils"
	"github.com/wahidx/kafkax/xkafka"
)

// topicCmd represents the topic command
var createTopicCmd = &cobra.Command{
	Use:   "topic",
	Short: "create topic",

	Run: func(cmd *cobra.Command, args []string) {
		broker, err := cmd.Flags().GetString("broker")
		if err != nil {
			fmt.Println("Invalid broker")
			return
		} else if len(broker) != 0 {
			utils.SetBroker(broker)
		}

		topic, err := cmd.Flags().GetString("topic")
		if err != nil {
			fmt.Println("Invalid topic name")
			return
		}

		partition, err := cmd.Flags().GetInt("partition")
		if err != nil {
			fmt.Println("Invalid partition number")
			return
		}

		replica, err := cmd.Flags().GetInt("replica")
		if err != nil {
			fmt.Println("Invalid number of replica")
			return
		}

		xkafka.CreateTopic(topic, partition, replica)
	},
}

func init() {
	createCmd.AddCommand(createTopicCmd)

	createCmd.PersistentFlags().StringP("topic", "t", "", "Topic name")
	createCmd.PersistentFlags().IntP("partition", "p", 1, "Partition number")
	createCmd.PersistentFlags().IntP("replica", "r", 1, "Number of replica")
}
