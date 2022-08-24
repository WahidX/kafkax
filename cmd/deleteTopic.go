/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wahidx/kafkax/xkafka"
)

// deleteTopicCmd represents the deleteTopic command
var deleteTopicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Delete topics",
	Long:  "For multiple topics, pass comma separated topics",

	Run: func(cmd *cobra.Command, args []string) {
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
