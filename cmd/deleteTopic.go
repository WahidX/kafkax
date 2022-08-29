/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wahidx/kafkax/utils"
	"github.com/wahidx/kafkax/xkafka"
)

// deleteTopicCmd represents the deleteTopic command
var deleteTopicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Delete topics",
	Long:  "For multiple topics, pass comma separated topics",

	Run: func(cmd *cobra.Command, args []string) {
		broker, err := cmd.Flags().GetString("broker")
		if err != nil {
			fmt.Println("Invalid broker")
			return
		} else if len(broker) != 0 {
			utils.SetBroker(broker)
		}

		topicsStr, err := cmd.Flags().GetString("topic")
		if err != nil || len(topicsStr) == 0 {
			fmt.Println("Invalid topics")
			return
		}

		topics := strings.Split(topicsStr, ",")
		xkafka.DeleteTopic(topics...)
	},
}

func init() {
	deleteCmd.AddCommand(deleteTopicCmd)

	deleteCmd.PersistentFlags().StringP("topic", "t", "", "Topic names")
}
