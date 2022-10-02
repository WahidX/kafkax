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
var listTopicCmd = &cobra.Command{
	Use:   "topic",
	Short: "List all topics",

	Run: func(cmd *cobra.Command, args []string) {
		broker, err := cmd.Flags().GetString("broker")
		if err != nil {
			fmt.Println("Invalid broker")
			return
		} else if len(broker) != 0 {
			utils.SetBroker(broker)
		}

		j, err := cmd.Flags().GetBool("json")
		if err != nil {
			fmt.Println("Invalid search key")
			return
		}

		key, err := cmd.Flags().GetString("sKey")
		if err != nil {
			fmt.Println("Invalid search key")
			return
		} else if len(key) > 0 {
			xkafka.FindTopics(key, j)
		} else {
			xkafka.ListTopics(j)
		}

	},
}

func init() {
	listCmd.AddCommand(listTopicCmd)

	listCmd.PersistentFlags().StringP("sKey", "s", "", "Search key")
	listCmd.PersistentFlags().BoolP("json", "j", false, "JSON stdout")
}
